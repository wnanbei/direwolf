package direwolf

import (
	"net/http"
	"net/url"
)

// Headers is request headers, as parameter in Request method.
// Headers type is http.Header, so you can init it like this:
// 	headers := Headers{
// 		"key1": {"value1", "value2"},
// 		"key2": {"value3"},
//  }
type Headers http.Header

// Cookies is request cookies, as parameter in Request method.
type Cookies map[string]string

// DataForm is the form you want to post, as parameter in Request method.
// DataForm type is url.values, so you can init it like this:
// 	df := DataForm{
// 		"key1": {"value1", "value2"},
// 		"key2": {"value3"},
//  }
type DataForm url.Values

// Data is data you want to post, as parameter in Request method.
type Data string

// Params is url params you want to join to url, as parameter in Request method.
// Params type is url.values, so you can init it like this:
// 	params := Params{
//  	"key1": {"value1", "value2"},
//  	"key2": {"value3"},
//  }
type Params url.Values

// Get is the most common method of direwolf to initiate a get request.
// @url: request url, necessary.
func Get(url string, args ...interface{}) {
	session := Session{}
	session.Get(url, args...)
}

// Post is the method to initiate a post request.
// @url: request url, necessary.
func Post(url string, args ...interface{}) {
	session := Session{}
	session.Post(url, args...)
}
