package direwolf

// Get is the most common method of direwolf to initiate a get request.
// @url: request url, necessary.
func Get(url string, args ...interface{}) *Response {
	session := NewSession()
	return session.Get(url, args...)
}

// Post is the method to initiate a post request.
// @url: request url, necessary.
func Post(url string, args ...interface{}) *Response {
	session := NewSession()
	return session.Post(url, args...)
}
