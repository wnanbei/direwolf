package direwolf

import (
	"net/http"
	"net/url"
	"sort"
	"strings"
)

// Params is url params you want to join to url, as parameter in Request method.
// You should init it by using NewParams like this:
// 	params := dw.NewParams(
//		"key1", "value1",
// 		"key2", "value2",
// 	)
// Note: mid symbol is comma.
type Params struct {
	stringSliceMap
}

// Body is data you want to post, as parameter in Request method.
type Body []byte

// PostForm is the form you want to post, as parameter in Request method.
// You should init it by using NewPostForm like this:
// 	postform := dw.NewPostForm(
//		"key1", "value1",
// 		"key2", "value2",
// 	)
// Note: mid symbol is comma.
type PostForm struct {
	stringSliceMap
}

// Cookies is request cookies, as parameter in Request method.
// You should init it by using NewCookies like this:
// 	cookies := dw.NewCookies(
//		"key1", "value1",
// 		"key2", "value2",
// 	)
// Note: mid symbol is comma.
type Cookies struct {
	stringSliceMap
}

// Proxy is the proxy server address, like "http://127.0.0.1:1080".
type Proxy string

// RedirectNum is the number of request redirect allowed.
type RedirectNum int

// Timeout is the number of time to timeout request.
// if timeout > 0, it means a time limit for requests.
// if timeout < 0, it means no limit.
// if timeout = 0, it means keep default 30 second timeout.
type Timeout int

// stringSliceMap type is map[string][]string, used for Params, PostForm, Cookies.
type stringSliceMap struct {
	data map[string][]string
}

// New is the way to create a stringSliceMap.
// You can set key-value pair when you init it by sent params. Just like this:
// 	stringSliceMap{}.New(
// 		"key1", "value1",
// 		"key2", "value2",
// 	)
// But be careful, between the key and value is a comma.
// And if the number of parameters is not a multiple of 2, it will panic.
func (ssm *stringSliceMap) New(keyvalue ...string) {
	ssm.data = make(map[string][]string)
	if keyvalue != nil {
		if len(keyvalue)%2 != 0 {
			panic("key and value must be pair")
		}

		for i := 0; i < len(keyvalue)/2; i++ {
			key := keyvalue[i*2]
			value := keyvalue[i*2+1]
			ssm.data[key] = append(ssm.data[key], value)
		}
	}
}

// Add key and value to stringSliceMap.
// If key exiests, value will append to slice.
func (ssm *stringSliceMap) Add(key, value string) {
	ssm.data[key] = append(ssm.data[key], value)
}

// Set key and value to stringSliceMap.
// If key exiests, existed value will drop and new value will set.
func (ssm *stringSliceMap) Set(key, value string) {
	ssm.data[key] = []string{value}
}

// Del delete the given key.
func (ssm *stringSliceMap) Del(key string) {
	delete(ssm.data, key)
}

// Get get the value pair to given key.
// You can pass index to assign which value to get, when there are multiple values.
func (ssm *stringSliceMap) Get(key string, index ...int) string {
	if ssm.data == nil {
		return ""
	}
	ssmValue := ssm.data[key]
	if len(ssmValue) == 0 {
		return ""
	}
	if index != nil {
		return ssmValue[index[0]]
	}
	return ssmValue[0]
}

// URLEncode encodes the values into ``URL encoded'' form
// ("bar=baz&foo=quux") sorted by key.
func (ssm *stringSliceMap) URLEncode() string {
	if ssm.data == nil {
		return ""
	}
	var buf strings.Builder
	keys := make([]string, 0, len(ssm.data))
	for k := range ssm.data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		ssmValue := ssm.data[k]
		keyEscaped := url.QueryEscape(k)
		for _, v := range ssmValue {
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(keyEscaped)
			buf.WriteByte('=')
			buf.WriteString(url.QueryEscape(v))
		}
	}
	return buf.String()
}

// NewParams new a Params type.
//
// You can set key-value pair when you init it by sent parameters. Just like this:
// 	params := NewParams(
// 		"key1", "value1",
// 		"key2", "value2",
// 	)
// But be careful, between the key and value is a comma.
// And if the number of parameters is not a multiple of 2, it will panic.
func NewParams(keyvalue ...string) *Params {
	var p = &Params{}
	p.New(keyvalue...)
	return p
}

// NewCookies new a Cookies type.
//
// You can set key-value pair when you init it by sent parameters. Just like this:
// 	cookies := NewCookies(
// 		"key1", "value1",
// 		"key2", "value2",
// 	)
// But be careful, between the key and value is a comma.
// And if the number of parameters is not a multiple of 2, it will panic.
func NewCookies(keyvalue ...string) *Cookies {
	var c = &Cookies{}
	c.New(keyvalue...)
	return c
}

// NewPostForm new a PostForm type.
//
// You can set key-value pair when you init it by sent parameters. Just like this:
// 	postform := NewPostForm(
// 		"key1", "value1",
// 		"key2", "value2",
// 	)
// But be careful, between the key and value is a comma.
// And if the number of parameters is not a multiple of 2, it will panic.
func NewPostForm(keyvalue ...string) *PostForm {
	var p = &PostForm{}
	p.New(keyvalue...)
	return p
}

// NewHeaders new a http.Header type.
//
// You can set key-value pair when you init it by sent parameters. Just like this:
// 	headers := NewHeaders(
// 		"key1", "value1",
// 		"key2", "value2",
// 	)
// But be careful, between the key and value is a comma.
// And if the number of parameters is not a multiple of 2, it will panic.
func NewHeaders(keyvalue ...string) http.Header {
	h := http.Header{}
	if keyvalue != nil {
		if len(keyvalue)%2 != 0 {
			panic("key and value must be part")
		}

		for i := 0; i < len(keyvalue)/2; i++ {
			key := keyvalue[i*2]
			value := keyvalue[i*2+1]
			h.Add(key, value)
		}
	}
	return h
}
