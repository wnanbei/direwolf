package direwolf

import (
	"net/http"
	"net/url"
	"sort"
	"strings"
)

// Params is url params you want to join to url, as parameter in Request method.
// You can init it like this:
// 	params := Params{
//  	"key1": {"value1", "value2"},
//  	"key2": {"value3"},
//  }
type Params url.Values

// Headers is request headers, as parameter in Request method.
// You can init it like this:
// 	headers := Headers{
// 		"key1": {"value1", "value2"},
// 		"key2": {"value3"},
//  }
type Headers http.Header

// Data is data you want to post, as parameter in Request method.
type Data string

// DataForm is the form you want to post, as parameter in Request method.
// You can init it like this:
// 	df := DataForm{
// 		"key1": {"value1", "value2"},
// 		"key2": {"value3"},
//  }
type DataForm url.Values

// Cookies is request cookies, as parameter in Request method.
// You can init it like this:
// 	c := Cookies{
// 		"key1": "value1",
// 		"key2": "value2",
//  }
type Cookies map[string]string

// Proxy is the proxy server address, like "http://127.0.0.1:1080"
type Proxy string

// RedirectNum is the number of redirect
type RedirectNum int

// Timeout is the number of time to timeout request.
type Timeout int

// StringSliceMap type is map[string][]string, used for Params, PostForms, Cookies.
// you should new it like this:
//   StringSliceMap{}.New()
type StringSliceMap map[string][]string

// New is the most convenient way to create a StringSliceMap.
// You can set key-value pair when you init it by sent params. Just like this:
// StringSliceMap{}.New(
// 	"key1", "value1",
// 	"key2", "value2",
// )
// But be careful, between the key and value is a comma.
// And if the number of parameters is not a multiple of 2, it will panic.
func (ssm StringSliceMap) New(keyvalue ...string) StringSliceMap {
	ssm = make(map[string][]string)
	if keyvalue != nil {
		if len(keyvalue)%2 != 0 {
			panic("key and value must be part")
		}

		for i := 0; i < len(keyvalue)/2; i++ {
			key := keyvalue[i*2]
			value := keyvalue[i*2+1]
			ssm[key] = append(ssm[key], value)
		}
	}
	return ssm
}

// Add key and value to StringSliceMap.
// If key exiests, value will append to slice.
func (ssm StringSliceMap) Add(key, value string) {
	ssm[key] = append(ssm[key], value)
}

// Set key and value to StringSliceMap.
// If key exiests, existed value will drop and new value will set.
func (ssm StringSliceMap) Set(key, value string) {
	ssm[key] = []string{value}
}

// Del delete the given key.
func (ssm StringSliceMap) Del(key string) {
	delete(ssm, key)
}

// Get get the value pair to given key.
// You can pass index to assign which value to get, when there are multiple values.
func (ssm StringSliceMap) Get(key string, index ...int) string {
	if ssm == nil {
		return ""
	}
	ssmValue := ssm[key]
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
func (ssm StringSliceMap) URLEncode() string {
	if ssm == nil {
		return ""
	}
	var buf strings.Builder
	keys := make([]string, 0, len(ssm))
	for k := range ssm {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		ssmValue := ssm[k]
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
