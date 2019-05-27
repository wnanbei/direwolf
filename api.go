package direwolf

// Get is the most common method of direwolf to initiate a get request.
// @url: request url, necessary.
func Get(url string, args ...interface{}) *Response {
	session := Session{}
	return session.Get(url, args...)
}

// Post is the method to initiate a post request.
// @url: request url, necessary.
func Post(url string, args ...interface{}) *Response {
	session := Session{}
	return session.Post(url, args...)
}

// Headers is request headers, as parameter in Request method.
// You can init it like this:
// 	headers := Headers{
// 		"key1": {"value1", "value2"},
// 		"key2": {"value3"},
//  }
type Headers map[string][]string

// Params is url params you want to join to url, as parameter in Request method.
// You can init it like this:
// 	params := Params{
//  	"key1": {"value1", "value2"},
//  	"key2": {"value3"},
//  }
type Params map[string][]string

// DataForm is the form you want to post, as parameter in Request method.
// You can init it like this:
// 	df := DataForm{
// 		"key1": {"value1", "value2"},
// 		"key2": {"value3"},
//  }
type DataForm map[string][]string

// Cookies is request cookies, as parameter in Request method.
// You can init it like this:
// 	c := Cookies{
// 		"key1": "value1",
// 		"key2": "value2",
//  }
type Cookies map[string]string

// Data is data you want to post, as parameter in Request method.
type Data string
