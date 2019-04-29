package direwolf

// Headers is request headers, as parameter in Request method.
type Headers map[string]string

// Cookies is request cookies, as parameter in Request method.
type Cookies map[string]string

// DataForm is the form you want to post, as parameter in Request method.
type DataForm map[string]string

// Data is data you want to post, as parameter in Request method.
type Data string

// Params is url params you want to join to url, as parameter in Request method.
type Params map[string]string

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
