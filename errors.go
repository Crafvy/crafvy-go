package crafvy

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

// Sentinels to test an *Error against with errors.Is:
//
//	if errors.Is(err, crafvy.ErrNotFound) { … }
var (
	ErrValidation     = errors.New("crafvy: invalid request (400)")
	ErrAuthentication = errors.New("crafvy: no valid token (401)")
	ErrAuthorization  = errors.New("crafvy: not allowed (403)")
	ErrNotFound       = errors.New("crafvy: not found (404)")
	ErrConflict       = errors.New("crafvy: conflict (409)")
	ErrRateLimit      = errors.New("crafvy: too many requests (429)")
	// ErrNoToken: NewClient got no token and CRAFVY_TOKEN is not set.
	ErrNoToken = errors.New("crafvy: no token — pass one to NewClient or set CRAFVY_TOKEN")
)

// Error is any error answer of the API (and a timeout, as status 408).
//
// Every refusal the API sends is {"success": false, "error": "…", "code": "…"}.
// Code is the API's machine-readable reason (tokenPermissionMissing,
// crmSpaceNotFound, …) — branch on it, not on the English Message:
//
//	var apiErr *crafvy.Error
//	if errors.As(err, &apiErr) && apiErr.Code == "tasksSpaceRequired" { … }
type Error struct {
	// Status is the HTTP status.
	Status int
	// Code is the API's stable reason, when it sent one.
	Code string
	// Message is the API's English sentence, or "HTTP <status>".
	Message string
	// Body is the answer as it came.
	Body json.RawMessage
	// RetryAfter is what a 429 said to wait (after the SDK's own retries); 0 when it did not say.
	RetryAfter time.Duration

	err error
}

func (e *Error) Error() string {
	if e.Code != "" {
		return fmt.Sprintf("crafvy: %d %s: %s", e.Status, e.Code, e.Message)
	}
	return fmt.Sprintf("crafvy: %d: %s", e.Status, e.Message)
}

// Is matches the sentinel of the error's status.
func (e *Error) Is(target error) bool {
	switch target {
	case ErrValidation:
		return e.Status == 400
	case ErrAuthentication:
		return e.Status == 401
	case ErrAuthorization:
		return e.Status == 403
	case ErrNotFound:
		return e.Status == 404
	case ErrConflict:
		return e.Status == 409
	case ErrRateLimit:
		return e.Status == 429
	}
	return false
}

// Unwrap returns the underlying error of a timeout.
func (e *Error) Unwrap() error { return e.err }

func newError(status int, body []byte, retryAfter time.Duration) *Error {
	e := &Error{Status: status, Message: fmt.Sprintf("HTTP %d", status)}
	if len(body) > 0 {
		if json.Valid(body) {
			e.Body = json.RawMessage(body)
		} else {
			e.Body, _ = json.Marshal(string(body))
		}
	}
	var parsed struct {
		Error any `json:"error"`
		Code  any `json:"code"`
	}
	if json.Unmarshal(body, &parsed) == nil {
		if s, ok := parsed.Error.(string); ok && s != "" {
			e.Message = s
		}
		if s, ok := parsed.Code.(string); ok {
			e.Code = s
		}
	}
	if status == 429 && retryAfter > 0 {
		e.RetryAfter = retryAfter
	}
	return e
}
