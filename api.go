package direwolf

// Get is the most common method of direwolf to initiate a get request.
// @url: request url, necessary.
func Get(url string, args ...interface{}) (*Response, error) {
	session := NewSession()
	resp, err := session.Get(url, args...)
	if err != nil {
		return nil, MakeErrorStack(err, "direwolf.Get()")
	}
	return resp, nil
}

// Post is the method to initiate a post request.
// @url: request url, necessary.
func Post(url string, args ...interface{}) (*Response, error) {
	session := NewSession()
	resp, err := session.Post(url, args...)
	if err != nil {
		return nil, MakeErrorStack(err, "direwolf.Post()")
	}
	return resp, nil
}
