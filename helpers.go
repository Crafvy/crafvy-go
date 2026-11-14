package crafvy

import (
	"os"
	"path/filepath"
)

// String returns a pointer to s — for optional fields: Name: crafvy.String("Anna").
func String(s string) *string { return &s }

// Int returns a pointer to i.
func Int(i int) *int { return &i }

// Float64 returns a pointer to f.
func Float64(f float64) *float64 { return &f }

// Bool returns a pointer to b — crafvy.Bool(false) is sent, unlike a plain false with omitempty.
func Bool(b bool) *bool { return &b }

// Ptr returns a pointer to any value.
func Ptr[T any](v T) *T { return &v }

// File is an upload for a multipart operation (assets.upload, …).
type File struct {
	// Name is the file name the API sees; the form field's name when empty.
	Name string
	// ContentType defaults to application/octet-stream.
	ContentType string
	Data        []byte
}

// FileFromPath reads a file from disk for an upload.
func FileFromPath(path, contentType string) (*File, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return &File{Name: filepath.Base(path), ContentType: contentType, Data: data}, nil
}
