package crafvy

import (
	"context"
	"errors"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"strings"
	"testing"
)

func testClient(rec *recorder, opts ...Option) *Client {
	return NewClient("pb_test", append([]Option{WithBaseURL("https://api.test"), WithHTTPClient(&http.Client{Transport: rec})}, opts...)...)
}

func TestPathParametersAreEncoded(t *testing.T) {
	rec := &recorder{reply: []byte(`{"success":true}`)}
	c := testClient(rec)
	if _, err := c.Projects.Get(context.Background(), "a b/c"); err != nil {
		t.Fatal(err)
	}
	if got := rec.reqs[0].URL.EscapedPath(); !strings.HasPrefix(got, "/api/projects/a%20b%2Fc") {
		t.Errorf("path %s", got)
	}
	if ua := rec.reqs[0].Header.Get("User-Agent"); ua != "crafvy-go/"+Version {
		t.Errorf("user agent %q", ua)
	}
}

func TestErrorsCarryStatusAndCode(t *testing.T) {
	rec := &recorder{status: 403, reply: []byte(`{"success":false,"error":"No","code":"tokenPermissionMissing"}`)}
	c := testClient(rec, WithRetries(0))
	_, err := c.Raw(context.Background(), "GET", "/api/projects", nil, nil)
	var apiErr *Error
	if !errors.As(err, &apiErr) {
		t.Fatalf("not an *Error: %v", err)
	}
	if apiErr.Status != 403 || apiErr.Code != "tokenPermissionMissing" || apiErr.Message != "No" {
		t.Errorf("%+v", apiErr)
	}
	if !errors.Is(err, ErrAuthorization) || errors.Is(err, ErrNotFound) {
		t.Error("sentinels")
	}

	rec.status, rec.reply = 404, []byte(`{"success":false,"error":"Not found"}`)
	_, err = c.Raw(context.Background(), "GET", "/api/projects/x", nil, nil)
	if !errors.Is(err, ErrNotFound) {
		t.Errorf("want ErrNotFound, got %v", err)
	}

	rec.status, rec.reply = 418, nil
	_, err = c.Raw(context.Background(), "GET", "/api/x", nil, nil)
	if !errors.As(err, &apiErr) || apiErr.Status != 418 || apiErr.Message != "HTTP 418" {
		t.Errorf("418: %v", err)
	}
}

func TestRetries429AndIdempotent5xxOnly(t *testing.T) {
	wait := http.Header{"Retry-After": []string{"0"}}
	rec := &recorder{replies: []reply{{429, wait, ""}, {200, nil, `{"ok":true}`}}}
	c := testClient(rec)
	out, err := c.Raw(context.Background(), "POST", "/api/x", nil, map[string]any{})
	if err != nil || string(out) != `{"ok":true}` || len(rec.reqs) != 2 {
		t.Fatalf("429 retried: %s %v (%d requests)", out, err, len(rec.reqs))
	}

	rec = &recorder{replies: []reply{{503, wait, `{"error":"down"}`}}, status: 200, reply: []byte(`{}`)}
	c = testClient(rec)
	_, err = c.Raw(context.Background(), "POST", "/api/x", nil, map[string]any{})
	if err == nil || len(rec.reqs) != 1 {
		t.Fatalf("a POST that failed with a 5xx is not sent again (%d requests, %v)", len(rec.reqs), err)
	}

	rec = &recorder{replies: []reply{{503, wait, ""}, {502, wait, ""}}, reply: []byte(`{"ok":1}`)}
	c = testClient(rec)
	if _, err := c.Raw(context.Background(), "GET", "/api/x", nil, nil); err != nil || len(rec.reqs) != 3 {
		t.Fatalf("a GET is retried on 5xx (%d requests, %v)", len(rec.reqs), err)
	}

	rec = &recorder{status: 429, header: wait, reply: []byte(`{"error":"slow down"}`)}
	c = testClient(rec, WithRetries(2))
	_, err = c.Raw(context.Background(), "GET", "/api/x", nil, nil)
	if !errors.Is(err, ErrRateLimit) || len(rec.reqs) != 3 {
		t.Fatalf("429 after retries: %v (%d requests)", err, len(rec.reqs))
	}
}

func TestStructuredQueryValuesTravelAsJSON(t *testing.T) {
	rec := &recorder{reply: []byte(`{}`)}
	c := testClient(rec)
	_, err := c.Raw(context.Background(), "GET", "/api/x", map[string]any{"filters": []any{map[string]any{"a": 1}}, "on": true, "skip": nil, "n": 2.5}, nil)
	if err != nil {
		t.Fatal(err)
	}
	q := rec.reqs[0].URL.Query()
	if q.Get("filters") != `[{"a":1}]` || q.Get("on") != "true" || q.Has("skip") || q.Get("n") != "2.5" {
		t.Errorf("query %v", q)
	}
}

func TestOptionalZeroValuesAreSentAndAbsentOnesAreNot(t *testing.T) {
	type body struct {
		Done  *bool    `json:"done,omitempty"`
		Tags  []string `json:"tags,omitempty"`
		Notes *string  `json:"notes,omitempty"`
		Name  string   `json:"name"`
		Lead  *string  `json:"lead"`
	}
	b, err := encodeJSON(&body{Done: Bool(false), Tags: []string{}, Name: "<a&b>"})
	if err != nil {
		t.Fatal(err)
	}
	if string(b) != `{"done":false,"tags":[],"name":"<a&b>","lead":null}` {
		t.Errorf("%s", b)
	}
}

func TestMultipartUploadsFiles(t *testing.T) {
	rec := &recorder{reply: []byte(`{"success":true}`)}
	c := testClient(rec)
	_, err := c.Assets.Upload(context.Background(), &AssetsUploadBody{
		File:           &File{Name: "a.txt", ContentType: "text/plain", Data: []byte("hello")},
		OrganizationID: "org",
		Visibility:     String("private"),
	})
	if err != nil {
		t.Fatal(err)
	}
	req := rec.reqs[0]
	mt, params, _ := mime.ParseMediaType(req.Header.Get("Content-Type"))
	if mt != "multipart/form-data" {
		t.Fatalf("content type %s", mt)
	}
	r := multipart.NewReader(strings.NewReader(string(rec.bodies[0])), params["boundary"])
	parts := map[string]string{}
	files := map[string]string{}
	for {
		p, err := r.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatal(err)
		}
		data, _ := io.ReadAll(p)
		parts[p.FormName()] = string(data)
		if p.FileName() != "" {
			files[p.FormName()] = p.FileName()
		}
	}
	if parts["file"] != "hello" || files["file"] != "a.txt" || parts["organizationId"] != "org" || parts["visibility"] != "private" {
		t.Errorf("parts %v files %v", parts, files)
	}
	if _, ok := parts["workspaceId"]; ok {
		t.Error("an absent field is not sent")
	}
}

func TestTokenFromEnvironment(t *testing.T) {
	t.Setenv("CRAFVY_TOKEN", "pb_env")
	rec := &recorder{reply: []byte(`{}`)}
	c := NewClient("", WithBaseURL("https://api.test"), WithHTTPClient(&http.Client{Transport: rec}))
	if _, err := c.Raw(context.Background(), "GET", "/api/x", nil, nil); err != nil {
		t.Fatal(err)
	}
	if got := rec.reqs[0].Header.Get("Authorization"); got != "Bearer pb_env" {
		t.Errorf("authorization %q", got)
	}
	t.Setenv("CRAFVY_TOKEN", "")
	if _, err := NewClient("").Raw(context.Background(), "GET", "/api/x", nil, nil); !errors.Is(err, ErrNoToken) {
		t.Errorf("want ErrNoToken, got %v", err)
	}
}

func TestReturnsPicksTheValue(t *testing.T) {
	rec := &recorder{reply: []byte(`{"success":true,"data":{"id":"p1","name":"Site","extra":1}}`)}
	c := testClient(rec)
	p, err := c.Projects.Get(context.Background(), "p1")
	if err != nil || p == nil || p.ID != "p1" {
		t.Fatalf("%+v %v", p, err)
	}
}

func TestCallOptions(t *testing.T) {
	rec := &recorder{reply: []byte(`{}`)}
	c := testClient(rec)
	if _, err := c.Raw(context.Background(), "get", "/api/x", nil, nil, CallHeader("X-Trace", "1")); err != nil {
		t.Fatal(err)
	}
	if rec.reqs[0].Header.Get("X-Trace") != "1" || rec.reqs[0].Method != "GET" {
		t.Errorf("%v %s", rec.reqs[0].Header, rec.reqs[0].Method)
	}
}
