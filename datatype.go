package direwolf

import (
	"net/http"
	"net/url"
	"sort"
	"strings"
)

// RequestOption is the interface of Request Options. Use to bind the options to Request.
type RequestOption interface {
	// Bind Options to Request
	bindRequest(request *Request) error
}

// Body is the data you want to post, one of the Request Options.
type Body []byte

// RequestOption interface method, bind request option to request.
func (options Body) bindRequest(request *Request) error {
	request.Body = options
	return nil
}

// RedirectNum is the number of request redirect allowed.
// If RedirectNum > 0, it means a redirect number limit for requests.
// If RedirectNum <= 0, it means ban redirect.
// If RedirectNum is not set, it means default 5 times redirect limit.
type RedirectNum int

// RequestOption interface method, bind request option to request.
func (options RedirectNum) bindRequest(request *Request) error {
	request.RedirectNum = int(options)
	return nil
}

// Timeout is the number of time to timeout request.
// if timeout > 0, it means a time limit for requests.
// if timeout < 0, it means no limit.
// if timeout = 0, it means keep default 30 second timeout.
type Timeout int

// RequestOption interface method, bind request option to request.
func (options Timeout) bindRequest(request *Request) error {
	request.Timeout = int(options)
	return nil
}

// Proxy is the proxy server address, like "http://127.0.0.1:1080".
// You can set different proxies for HTTP and HTTPS sites.
type Proxy struct {
	HTTP  string
	HTTPS string
}

// RequestOption interface method, bind request option to request.
func (options *Proxy) bindRequest(request *Request) error {
	request.Proxy = options
	return nil
}

// strSliceMap type is map[string][]string, used for Params, PostForm.
type strSliceMap struct {
	data map[string][]string
}

// New is the way to create a strSliceMap.
// You can set key-value pair when you init it by sent params. Just like this:
// 	stringSliceMap{}.New(
// 		"key1", "value1",
// 		"key2", "value2",
// 	)
// But be careful, between the key and value is a comma.
// And if the number of parameters is not a multiple of 2, it will panic.
func (ssm *strSliceMap) New(keyValue ...string) {
	ssm.data = make(map[string][]string)
	if keyValue != nil {
		if len(keyValue)%2 != 0 {
			panic("key and value must be pair")
		}

		for i := 0; i < len(keyValue)/2; i++ {
			key := keyValue[i*2]
			value := keyValue[i*2+1]
			ssm.data[key] = append(ssm.data[key], value)
		}
	}
}

// Add key and value to stringSliceMap.
// If key exists, value will append to slice.
func (ssm *strSliceMap) Add(key, value string) {
	ssm.data[key] = append(ssm.data[key], value)
}

// Set key and value to stringSliceMap.
// If key exists, existed value will drop and new value will set.
func (ssm *strSliceMap) Set(key, value string) {
	ssm.data[key] = []string{value}
}

// Del delete the given key.
func (ssm *strSliceMap) Del(key string) {
	delete(ssm.data, key)
}

// Get get the value pair to given key.
// You can pass index to assign which value to get, when there are multiple values.
func (ssm *strSliceMap) Get(key string, index ...int) string {
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
// ("bar=baz&foo=qux") sorted by key.
func (ssm *strSliceMap) URLEncode() string {
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

// Params is url params you want to join to url, as parameter in Request method.
// You should init it by using NewParams like this:
// 	params := dw.NewParams(
//		"key1", "value1",
// 		"key2", "value2",
// 	)
// Note: mid symbol is comma.
type Params struct {
	strSliceMap
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
func NewParams(keyValue ...string) *Params {
	var p = &Params{}
	p.New(keyValue...)
	return p
}

// RequestOption interface method, bind request option to request.
func (options *Params) bindRequest(request *Request) error {
	request.Params = options
	u, err := url.Parse(request.URL)
	if err != nil {
		return WrapErrf(err, "URL error")
	}

	spacer := "?"
	// check whether parameters is existed in url.
	if u.RawQuery != "" {
		spacer = "&"
	} else if u.ForceQuery == true {
		spacer = ""
	}
	request.URL = request.URL + spacer + request.Params.URLEncode()
	return nil
}

// PostForm is the form you want to post, as parameter in Request method.
// You should init it by using NewPostForm like this:
// 	postForm := dw.NewPostForm(
//		"key1", "value1",
// 		"key2", "value2",
// 	)
// Note: mid symbol is comma.
type PostForm struct {
	strSliceMap
}

// NewPostForm new a PostForm type.
//
// You can set key-value pair when you init it by sent parameters. Just like this:
// 	postForm := NewPostForm(
// 		"key1", "value1",
// 		"key2", "value2",
// 	)
// But be careful, between the key and value is a comma.
// And if the number of parameters is not a multiple of 2, it will panic.
func NewPostForm(keyValue ...string) *PostForm {
	var p = &PostForm{}
	p.New(keyValue...)
	return p
}

// RequestOption interface method, bind request option to request.
func (options *PostForm) bindRequest(request *Request) error {
	request.PostForm = options
	return nil
}

type Headers struct {
	http.Header
}

// RequestOption interface method, bind request option to request.
func (options Headers) bindRequest(request *Request) error {
	request.Headers = http.Header(options.Header)
	return nil
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
func NewHeaders(keyValue ...string) *Headers {
	h := new(Headers)
	h.Header = http.Header{}
	if keyValue != nil {
		if len(keyValue)%2 != 0 {
			panic("key and value must be part")
		}

		for i := 0; i < len(keyValue)/2; i++ {
			key := keyValue[i*2]
			value := keyValue[i*2+1]
			h.Add(key, value)
		}
	}
	return h
}

// Cookies is request cookies, as parameter in Request method.
// You should init it by using NewCookies like this:
// 	cookies := dw.NewCookies(
//		"key1", "value1",
// 		"key2", "value2",
// 	)
// Note: mid symbol is comma.
type Cookies []*http.Cookie

// NewCookies new a Cookies type.
//
// You can set key-value pair when you init it by sent parameters. Just like this:
// 	cookies := NewCookies(
// 		"key1", "value1",
// 		"key2", "value2",
// 	)
// But be careful, between the key and value is a comma.
// And if the number of parameters is not a multiple of 2, it will panic.
func NewCookies(keyValue ...string) Cookies {
	c := make(Cookies, 0)
	if keyValue != nil {
		if len(keyValue)%2 != 0 {
			panic("key and value must be part")
		}

		for i := 0; i < len(keyValue)/2; i++ {
			key := keyValue[i*2]
			value := keyValue[i*2+1]
			cookie := &http.Cookie{Name: key, Value: value}
			c = append(c, cookie)
		}
	}
	return c
}

// RequestOption interface method, bind request option to request.
func (c Cookies) bindRequest(request *Request) error {
	request.Cookies = c
	return nil
}

// Add append a new cookie to Cookies.
func (c Cookies) Add(key, value string) {
	c = append(c, &http.Cookie{Name: key, Value: value})
}
