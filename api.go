package direwolf

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
