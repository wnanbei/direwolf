package direwolf

import (
	"net/url"
)

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
