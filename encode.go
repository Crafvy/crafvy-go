package crafvy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/textproto"
	"net/url"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

// The request side has its own JSON encoder for one reason: `omitempty`
// in encoding/json drops false, 0, "" and empty lists. Here, with
// `omitempty`, only an ABSENT value is left out — a nil pointer, slice, map
// or interface — so `Tags: []string{}` clears a list and `Done: crafvy.Bool(false)`
// is sent. Without `omitempty` (a required field) nil is sent as null.

var marshalerType = reflect.TypeOf((*json.Marshaler)(nil)).Elem()

// encodeJSON encodes a request body.
func encodeJSON(v any) ([]byte, error) {
	var buf bytes.Buffer
	if err := writeJSON(&buf, reflect.ValueOf(v)); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func isNil(v any) bool {
	if v == nil {
		return true
	}
	return nilValue(reflect.ValueOf(v))
}

func nilValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Ptr, reflect.Interface, reflect.Map, reflect.Slice, reflect.Func, reflect.Chan:
		return v.IsNil()
	}
	return !v.IsValid()
}

func writeJSON(buf *bytes.Buffer, v reflect.Value) error {
	if !v.IsValid() || nilValue(v) {
		buf.WriteString("null")
		return nil
	}
	if v.Type().Implements(marshalerType) {
		b, err := v.Interface().(json.Marshaler).MarshalJSON()
		if err != nil {
			return err
		}
		return json.Compact(buf, b)
	}
	switch v.Kind() {
	case reflect.Ptr, reflect.Interface:
		return writeJSON(buf, v.Elem())
	case reflect.Struct:
		buf.WriteByte('{')
		first := true
		for _, f := range structFields(v.Type()) {
			fv := v.Field(f.index)
			if f.omitEmpty && nilValue(fv) {
				continue
			}
			if !first {
				buf.WriteByte(',')
			}
			first = false
			writeLeaf(buf, f.name)
			buf.WriteByte(':')
			if err := writeJSON(buf, fv); err != nil {
				return err
			}
		}
		buf.WriteByte('}')
		return nil
	case reflect.Map:
		if v.Type().Key().Kind() != reflect.String {
			return writeLeafValue(buf, v)
		}
		keys := v.MapKeys()
		sort.Slice(keys, func(i, j int) bool { return keys[i].String() < keys[j].String() })
		buf.WriteByte('{')
		for i, k := range keys {
			if i > 0 {
				buf.WriteByte(',')
			}
			writeLeaf(buf, k.String())
			buf.WriteByte(':')
			if err := writeJSON(buf, v.MapIndex(k)); err != nil {
				return err
			}
		}
		buf.WriteByte('}')
		return nil
	case reflect.Slice, reflect.Array:
		if v.Type().Elem().Kind() == reflect.Uint8 {
			return writeLeafValue(buf, v) // []byte: base64, as encoding/json does
		}
		buf.WriteByte('[')
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buf.WriteByte(',')
			}
			if err := writeJSON(buf, v.Index(i)); err != nil {
				return err
			}
		}
		buf.WriteByte(']')
		return nil
	}
	return writeLeafValue(buf, v)
}

func writeLeaf(buf *bytes.Buffer, v any) {
	_ = writeLeafValue(buf, reflect.ValueOf(v))
}

// writeLeafValue: encoding/json without its HTML escaping (<, >, & stay as they are).
func writeLeafValue(buf *bytes.Buffer, v reflect.Value) error {
	var tmp bytes.Buffer
	enc := json.NewEncoder(&tmp)
	enc.SetEscapeHTML(false)
	if err := enc.Encode(v.Interface()); err != nil {
		return err
	}
	buf.Write(bytes.TrimRight(tmp.Bytes(), "\n"))
	return nil
}

type fieldInfo struct {
	index     int
	name      string
	omitEmpty bool
}

func structFields(t reflect.Type) []fieldInfo {
	var out []fieldInfo
	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)
		if !sf.IsExported() {
			continue
		}
		tag := sf.Tag.Get("json")
		if tag == "-" {
			continue
		}
		name, opts, _ := strings.Cut(tag, ",")
		if name == "" {
			name = sf.Name
		}
		out = append(out, fieldInfo{index: i, name: name, omitEmpty: strings.Contains(","+opts+",", ",omitempty,")})
	}
	return out
}

// encodeQuery turns a query struct, map[string]any or url.Values into
// url.Values: absent values left out, booleans true/false, a structured
// value (object, list) as JSON text.
func encodeQuery(q any) (url.Values, error) {
	if isNil(q) {
		return nil, nil
	}
	if vals, ok := q.(url.Values); ok {
		return vals, nil
	}
	out := url.Values{}
	v := reflect.ValueOf(q)
	for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	switch v.Kind() {
	case reflect.Struct:
		for _, f := range structFields(v.Type()) {
			if err := addQuery(out, f.name, v.Field(f.index)); err != nil {
				return nil, err
			}
		}
	case reflect.Map:
		for _, k := range v.MapKeys() {
			if err := addQuery(out, fmt.Sprint(k.Interface()), v.MapIndex(k)); err != nil {
				return nil, err
			}
		}
	default:
		return nil, fmt.Errorf("crafvy: a query must be a struct, a map or url.Values, not %s", v.Type())
	}
	return out, nil
}

func addQuery(out url.Values, key string, v reflect.Value) error {
	s, ok, err := scalarText(v)
	if err != nil || !ok {
		return err
	}
	out.Set(key, s)
	return nil
}

// scalarText is how a value travels as text (a query value, a form field).
func scalarText(v reflect.Value) (string, bool, error) {
	for v.IsValid() && (v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface) {
		if v.IsNil() {
			return "", false, nil
		}
		v = v.Elem()
	}
	if !v.IsValid() {
		return "", false, nil
	}
	switch v.Kind() {
	case reflect.String:
		return v.String(), true, nil
	case reflect.Bool:
		return strconv.FormatBool(v.Bool()), true, nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10), true, nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(v.Uint(), 10), true, nil
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'f', -1, 64), true, nil
	case reflect.Map, reflect.Slice:
		if v.IsNil() {
			return "", false, nil
		}
	}
	b, err := encodeJSON(v.Interface())
	if err != nil {
		return "", false, err
	}
	return string(b), true, nil
}

// encodeMultipart writes a body struct (or map) as multipart/form-data:
// a *File or []byte as a file part, everything else as a text field.
func encodeMultipart(body any) ([]byte, string, error) {
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)
	v := reflect.ValueOf(body)
	for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	write := func(name string, fv reflect.Value) error {
		for fv.IsValid() && fv.Kind() == reflect.Interface && !fv.IsNil() {
			fv = fv.Elem()
		}
		if !fv.IsValid() || nilValue(fv) {
			return nil
		}
		switch x := fv.Interface().(type) {
		case *File:
			return writeFilePart(w, name, x)
		case File:
			return writeFilePart(w, name, &x)
		case []byte:
			return writeFilePart(w, name, &File{Name: name, Data: x})
		case []*File:
			for _, f := range x {
				if f != nil {
					if err := writeFilePart(w, name, f); err != nil {
						return err
					}
				}
			}
			return nil
		}
		s, ok, err := scalarText(fv)
		if err != nil || !ok {
			return err
		}
		return w.WriteField(name, s)
	}
	switch v.Kind() {
	case reflect.Struct:
		for _, f := range structFields(v.Type()) {
			if err := write(f.name, v.Field(f.index)); err != nil {
				return nil, "", err
			}
		}
	case reflect.Map:
		keys := v.MapKeys()
		sort.Slice(keys, func(i, j int) bool { return fmt.Sprint(keys[i]) < fmt.Sprint(keys[j]) })
		for _, k := range keys {
			if err := write(fmt.Sprint(k.Interface()), v.MapIndex(k)); err != nil {
				return nil, "", err
			}
		}
	default:
		return nil, "", fmt.Errorf("crafvy: a multipart body must be a struct or a map, not %s", v.Type())
	}
	if err := w.Close(); err != nil {
		return nil, "", err
	}
	return buf.Bytes(), w.FormDataContentType(), nil
}

func writeFilePart(w *multipart.Writer, field string, f *File) error {
	name := f.Name
	if name == "" {
		name = field
	}
	ct := f.ContentType
	if ct == "" {
		ct = "application/octet-stream"
	}
	h := textproto.MIMEHeader{}
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, escapeQuotes(field), escapeQuotes(name)))
	h.Set("Content-Type", ct)
	part, err := w.CreatePart(h)
	if err != nil {
		return err
	}
	_, err = part.Write(f.Data)
	return err
}

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string { return quoteEscaper.Replace(s) }
