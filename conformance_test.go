package crafvy

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"reflect"
	"testing"
)

// The shared suite: in the Crafvy repository sdks/conformance.json, in the
// public SDK repository beside the package.
func loadSuite(t *testing.T) suite {
	t.Helper()
	for _, p := range []string{"conformance.json", "../conformance.json"} {
		data, err := os.ReadFile(p)
		if err != nil {
			continue
		}
		var s suite
		if err := json.Unmarshal(data, &s); err != nil {
			t.Fatal(err)
		}
		return s
	}
	t.Fatal("conformance.json not found")
	return suite{}
}

type suite struct {
	Version string `json:"version"`
	Cases   []struct {
		ID       string          `json:"id"`
		PathArgs []string        `json:"pathArgs"`
		Query    json.RawMessage `json:"query"`
		Body     json.RawMessage `json:"body"`
		Expect   struct {
			Method string            `json:"method"`
			Path   string            `json:"path"`
			Query  map[string]string `json:"query"`
			Body   json.RawMessage   `json:"body"`
		} `json:"expect"`
		Response json.RawMessage `json:"response"`
		Result   json.RawMessage `json:"result"`
	} `json:"cases"`
}

// recorder answers every request with `reply` and keeps what was sent.
type recorder struct {
	reqs   []*http.Request
	bodies [][]byte
	status int
	header http.Header
	reply  []byte
	// replies, when set, are used one per request before `reply`.
	replies []reply
}

type reply struct {
	status int
	header http.Header
	body   string
}

func (r *recorder) RoundTrip(req *http.Request) (*http.Response, error) {
	var body []byte
	if req.Body != nil {
		body, _ = io.ReadAll(req.Body)
	}
	r.reqs = append(r.reqs, req)
	r.bodies = append(r.bodies, body)
	status, header, data := r.status, r.header, r.reply
	if len(r.replies) > 0 {
		next := r.replies[0]
		r.replies = r.replies[1:]
		status, header, data = next.status, next.header, []byte(next.body)
	}
	if status == 0 {
		status = 200
	}
	if header == nil {
		header = http.Header{}
	}
	return &http.Response{StatusCode: status, Header: header, Body: io.NopCloser(bytes.NewReader(data)), Request: req}, nil
}

func decodeArg(raw json.RawMessage, into any) error {
	if len(raw) == 0 || string(raw) == "null" {
		return nil
	}
	return json.Unmarshal(raw, into)
}

// normalise decodes JSON and drops object keys whose value is null: a
// typed struct sends null for a required field the sample left out (a
// recursive shape cut short), which says the same thing.
func normalise(t *testing.T, data []byte) any {
	t.Helper()
	var v any
	if err := json.Unmarshal(data, &v); err != nil {
		t.Fatalf("not JSON: %s", data)
	}
	return dropNulls(v)
}

func dropNulls(v any) any {
	switch x := v.(type) {
	case map[string]any:
		for k, e := range x {
			if e == nil {
				delete(x, k)
			} else {
				x[k] = dropNulls(e)
			}
		}
	case []any:
		for i, e := range x {
			x[i] = dropNulls(e)
		}
	}
	return v
}

func TestConformance(t *testing.T) {
	s := loadSuite(t)
	calls := map[string]func(context.Context, *Client, []string, json.RawMessage, json.RawMessage) (any, error){}
	for _, c := range conformanceCalls {
		calls[c.id] = c.call
	}
	rec := &recorder{}
	client := NewClient("pb_test", WithBaseURL("https://api.test"), WithHTTPClient(&http.Client{Transport: rec}), WithRetries(0))
	ran := 0
	for _, tc := range s.Cases {
		tc := tc
		t.Run(tc.ID, func(t *testing.T) {
			call, ok := calls[tc.ID]
			if !ok {
				t.Fatalf("no method for %s", tc.ID)
			}
			rec.reply = tc.Response
			result, err := call(context.Background(), client, tc.PathArgs, tc.Body, tc.Query)
			if err != nil {
				t.Fatal(err)
			}
			req := rec.reqs[len(rec.reqs)-1]
			sent := rec.bodies[len(rec.bodies)-1]
			if req.Method != tc.Expect.Method {
				t.Errorf("method %s, want %s", req.Method, tc.Expect.Method)
			}
			if req.URL.Path != tc.Expect.Path {
				t.Errorf("path %s, want %s", req.URL.Path, tc.Expect.Path)
			}
			got := map[string]string{}
			for k, v := range req.URL.Query() {
				got[k] = v[0]
			}
			want := tc.Expect.Query
			if want == nil {
				want = map[string]string{}
			}
			if !reflect.DeepEqual(got, want) {
				t.Errorf("query %v, want %v", got, want)
			}
			if len(tc.Expect.Body) > 0 && string(tc.Expect.Body) != "null" {
				if !reflect.DeepEqual(normalise(t, sent), normalise(t, tc.Expect.Body)) {
					t.Errorf("body %s, want %s", sent, tc.Expect.Body)
				}
			}
			if req.Header.Get("Authorization") != "Bearer pb_test" {
				t.Errorf("authorization %q", req.Header.Get("Authorization"))
			}
			encoded, err := encodeJSON(result)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(normalise(t, encoded), normalise(t, tc.Result)) {
				t.Errorf("result %s\nwant   %s", encoded, tc.Result)
			}
		})
		ran++
	}
	if ran == 0 {
		t.Error("no cases")
	}
	t.Logf("%d conformance cases", ran)
}
