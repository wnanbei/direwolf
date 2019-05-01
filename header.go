package direwolf

import (
	"io"
	"net/http"
	"net/url"
)

// Headers is request headers, as parameter in Request method.
// Headers type is http.Header, so you can init it like this:
// 	headers := Headers{
// 		"key1": {"value1", "value2"},
// 		"key2": {"value3"},
//  }
type Headers map[string][]string

// Add adds the key, value pair to the header.
// It appends to any existing values associated with key.
// The key is case insensitive; it is canonicalized by
// CanonicalHeaderKey.
func (h Headers) Add(key, value string) {
	http.Header(h).Add(key, value)
}

// Set sets the header entries associated with key to the
// single element value. It replaces any existing values
// associated with key. The key is case insensitive; it is
// canonicalized by textproto.CanonicalMIMEHeaderKey.
// To use non-canonical keys, assign to the map directly.
func (h Headers) Set(key, value string) {
	http.Header(h).Set(key, value)
}

// Get gets the first value associated with the given key. If
// there are no values associated with the key, Get returns "".
// It is case insensitive; textproto.CanonicalMIMEHeaderKey is
// used to canonicalize the provided key. To access multiple
// values of a key, or to use non-canonical keys, access the
// map directly.
func (h Headers) Get(key string) string {
	return http.Header(h).Get(key)
}

// Del deletes the values associated with key.
// The key is case insensitive; it is canonicalized by
// CanonicalHeaderKey.
func (h Headers) Del(key string) {
	http.Header(h).Del(key)
}

// Write writes a header in wire format.
func (h Headers) Write(w io.Writer) error {
	return http.Header(h).Write(w)
}

// WriteSubset writes a header in wire format.
// If exclude is not nil, keys where exclude[key] == true are not written.
func (h Headers) WriteSubset(w io.Writer, exclude map[string]bool) error {
	return http.Header(h).WriteSubset(w, exclude)
}

// Params is url params you want to join to url, as parameter in Request method.
// Params type is url.values, so you can init it like this:
// 	params := Params{
//  	"key1": {"value1", "value2"},
//  	"key2": {"value3"},
//  }
type Params map[string][]string

// Get gets the first value associated with the given key.
// If there are no values associated with the key, Get returns
// the empty string. To access multiple values, use the map
// directly.
func (p Params) Get(key string) string {
	return url.Values(p).Get(key)
}

// Set sets the key to value. It replaces any existing
// values.
func (p Params) Set(key, value string) {
	url.Values(p).Set(key, value)
}

// Add adds the value to key. It appends to any existing
// values associated with key.
func (p Params) Add(key, value string) {
	url.Values(p).Add(key, value)
}

// Del deletes the values associated with key.
func (p Params) Del(key string) {
	url.Values(p).Del(key)
}

// Encode encodes the values into ``URL encoded'' form
// ("bar=baz&foo=quux") sorted by key.
func (p Params) Encode() string {
	return url.Values(p).Encode()
}
