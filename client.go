// Package crafvy is the official Go client of the Crafvy API.
//
//	client := crafvy.NewClient(os.Getenv("CRAFVY_TOKEN"))
//	projects, err := client.Projects.List(ctx, nil)
//
// Every resource and method is generated from the API's description — the
// same one the other SDKs and the reference at https://docs.crafvy.com/sdk
// come from. A token acts as its owner: their roles, workspaces, feature
// gates and plan limits apply exactly as in the dashboard.
package crafvy

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

// DefaultBaseURL is where the API answers unless WithBaseURL or CRAFVY_BASE_URL say otherwise.
const DefaultBaseURL = "https://app.crafvy.com"

// Client calls the Crafvy API. Its fields are the API's resources
// (client.Projects, client.CRM.Records, …); it is safe for concurrent use.
type Client struct {
	resources

	token      string
	baseURL    string
	httpClient *http.Client
	retries    int
	timeout    time.Duration
}

// Option configures a Client.
type Option func(*Client)

// WithBaseURL points the client at another instance (self-hosted, staging).
func WithBaseURL(u string) Option {
	return func(c *Client) { c.baseURL = strings.TrimRight(u, "/") }
}

// WithHTTPClient replaces the http.Client — for proxies, custom transports and tests.
func WithHTTPClient(h *http.Client) Option {
	return func(c *Client) { c.httpClient = h }
}

// WithRetries sets how often a 429 (or a 5xx of GET, HEAD, PUT, DELETE) is retried. Default 3.
func WithRetries(n int) Option {
	return func(c *Client) { c.retries = n }
}

// WithTimeout sets how long one call may take, retries included. Default 60 s; 0 means no limit.
func WithTimeout(d time.Duration) Option {
	return func(c *Client) { c.timeout = d }
}

// NewClient returns a client for the token. An empty token falls back to the
// environment variable CRAFVY_TOKEN; the base URL to CRAFVY_BASE_URL, then
// DefaultBaseURL. Without any token, every call fails with ErrNoToken.
func NewClient(token string, opts ...Option) *Client {
	if token == "" {
		token = os.Getenv("CRAFVY_TOKEN")
	}
	base := os.Getenv("CRAFVY_BASE_URL")
	if base == "" {
		base = DefaultBaseURL
	}
	c := &Client{
		token:      token,
		baseURL:    strings.TrimRight(base, "/"),
		httpClient: http.DefaultClient,
		retries:    3,
		timeout:    60 * time.Second,
	}
	for _, o := range opts {
		o(c)
	}
	c.resources.init(c)
	return c
}

// CallOption overrides a setting for one call.
type CallOption func(*callConfig)

type callConfig struct {
	timeout    time.Duration
	timeoutSet bool
	retries    int
	retriesSet bool
	headers    http.Header
}

// CallTimeout limits this call (retries included); 0 means no limit.
func CallTimeout(d time.Duration) CallOption {
	return func(c *callConfig) { c.timeout, c.timeoutSet = d, true }
}

// CallRetries sets how often this call is retried.
func CallRetries(n int) CallOption {
	return func(c *callConfig) { c.retries, c.retriesSet = n, true }
}

// CallHeader adds a header to this call.
func CallHeader(key, value string) CallOption {
	return func(c *callConfig) {
		if c.headers == nil {
			c.headers = http.Header{}
		}
		c.headers.Add(key, value)
	}
}

// Raw calls any path of the API directly — for an endpoint this SDK version
// has no method for yet. query may be nil, a map[string]any, url.Values or a
// query struct; body anything encoding/json can encode. The answer comes
// back as it was sent.
func (c *Client) Raw(ctx context.Context, method, path string, query, body any, opts ...CallOption) (json.RawMessage, error) {
	var out json.RawMessage
	err := c.call(ctx, strings.ToUpper(method), path, query, body, false, "", &out, opts)
	return out, err
}

// call is what every generated method runs: send, then decode the answer
// (or the part of it at the dotted path `returns`) into out.
func (c *Client) call(ctx context.Context, method, path string, query, body any, multipart bool, returns string, out any, opts []CallOption) error {
	raw, err := c.send(ctx, method, path, query, body, multipart, opts)
	if err != nil {
		return err
	}
	if returns != "" {
		raw = pick(raw, returns)
	}
	if out == nil || len(raw) == 0 || string(raw) == "null" {
		return nil
	}
	if rm, ok := out.(*json.RawMessage); ok {
		*rm = raw
		return nil
	}
	if err := json.Unmarshal(raw, out); err != nil {
		var typeErr *json.UnmarshalTypeError
		if errors.As(err, &typeErr) {
			// A field whose shape differs from the description stays empty;
			// the rest is decoded. Raw returns the answer exactly.
			return nil
		}
		return fmt.Errorf("crafvy: decoding the answer of %s %s: %w", method, path, err)
	}
	return nil
}

var idempotent = map[string]bool{"GET": true, "HEAD": true, "PUT": true, "DELETE": true}

func (c *Client) send(ctx context.Context, method, path string, query, body any, multipart bool, opts []CallOption) (json.RawMessage, error) {
	if c.token == "" {
		return nil, ErrNoToken
	}
	cfg := callConfig{timeout: c.timeout, retries: c.retries}
	for _, o := range opts {
		o(&cfg)
	}
	u := c.baseURL + path
	q, err := encodeQuery(query)
	if err != nil {
		return nil, err
	}
	if len(q) > 0 {
		u += "?" + q.Encode()
	}
	var payload []byte
	contentType := ""
	if !isNil(body) {
		if multipart {
			payload, contentType, err = encodeMultipart(body)
		} else {
			payload, err = encodeJSON(body)
			contentType = "application/json"
		}
		if err != nil {
			return nil, err
		}
	}
	if cfg.timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, cfg.timeout)
		defer cancel()
	}

	for attempt := 0; ; attempt++ {
		var reader io.Reader
		if payload != nil {
			reader = bytes.NewReader(payload)
		}
		req, err := http.NewRequestWithContext(ctx, method, u, reader)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Authorization", "Bearer "+c.token)
		req.Header.Set("Accept", "application/json")
		req.Header.Set("User-Agent", "crafvy-go/"+Version)
		if contentType != "" {
			req.Header.Set("Content-Type", contentType)
		}
		for k, vs := range cfg.headers {
			req.Header[k] = append([]string(nil), vs...)
		}
		res, err := c.httpClient.Do(req)
		if err != nil {
			if errors.Is(err, context.DeadlineExceeded) && cfg.timeout > 0 {
				return nil, &Error{Status: 408, Message: fmt.Sprintf("request timed out after %s", cfg.timeout), err: err}
			}
			return nil, fmt.Errorf("crafvy: %s %s: %w", method, path, err)
		}
		data, err := io.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			return nil, fmt.Errorf("crafvy: reading the answer of %s %s: %w", method, path, err)
		}
		retryAfter := parseRetryAfter(res.Header.Get("Retry-After"))
		retryable := res.StatusCode == 429 || (idempotent[method] && (res.StatusCode == 500 || res.StatusCode == 502 || res.StatusCode == 503 || res.StatusCode == 504))
		if retryable && attempt < cfg.retries {
			if err := sleep(ctx, delay(attempt, retryAfter)); err != nil {
				return nil, &Error{Status: 408, Message: "request timed out while waiting to retry", err: err}
			}
			continue
		}
		if res.StatusCode >= 400 {
			return nil, newError(res.StatusCode, data, retryAfter)
		}
		if len(data) > 0 && !json.Valid(data) {
			// not JSON (a CSV export, plain text): handed back as a JSON string
			data, _ = json.Marshal(string(data))
		}
		return data, nil
	}
}

// pick returns the value at a dotted path of an answer (an operation's `returns`).
func pick(raw json.RawMessage, path string) json.RawMessage {
	for _, key := range strings.Split(path, ".") {
		var obj map[string]json.RawMessage
		if err := json.Unmarshal(raw, &obj); err != nil || obj == nil {
			return nil
		}
		raw = obj[key]
	}
	return raw
}

// parseRetryAfter reads seconds or an HTTP date; -1 when absent.
func parseRetryAfter(v string) time.Duration {
	if v == "" {
		return -1
	}
	if s, err := strconv.ParseFloat(v, 64); err == nil && s >= 0 {
		return time.Duration(s * float64(time.Second))
	}
	if t, err := http.ParseTime(v); err == nil {
		if d := time.Until(t); d > 0 {
			return d
		}
		return 0
	}
	return -1
}

func delay(attempt int, retryAfter time.Duration) time.Duration {
	if retryAfter >= 0 {
		return retryAfter
	}
	return time.Duration(500*(1<<attempt))*time.Millisecond + time.Duration(rand.Intn(100))*time.Millisecond
}

func sleep(ctx context.Context, d time.Duration) error {
	if d <= 0 {
		return ctx.Err()
	}
	t := time.NewTimer(d)
	defer t.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-t.C:
		return nil
	}
}

// pathSeg URL-encodes a path parameter: everything but A–Z a–z 0–9 - . _ ~.
func pathSeg(s string) string {
	var b strings.Builder
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch >= 'A' && ch <= 'Z' || ch >= 'a' && ch <= 'z' || ch >= '0' && ch <= '9' || ch == '-' || ch == '.' || ch == '_' || ch == '~' {
			b.WriteByte(ch)
		} else {
			fmt.Fprintf(&b, "%%%02X", ch)
		}
	}
	return b.String()
}

// Operation is one entry of Operations.
type Operation struct {
	ID     string
	Method string
	Path   string
}

var warned sync.Map

// deprecated is called by a deprecated method; the // Deprecated: comment
// is what editors and staticcheck show, so it only notes the call once.
func deprecated(id string) {
	if _, seen := warned.LoadOrStore(id, true); seen {
		return
	}
	if os.Getenv("CRAFVY_DEPRECATION_WARNINGS") != "" {
		fmt.Fprintf(os.Stderr, "crafvy: %s is deprecated — see its documentation for what to use instead\n", id)
	}
}
