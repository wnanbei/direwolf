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

// Request is the method to initiate a request.
func Request(reqSetting *RequestSetting) (*Response, error) {
	session := NewSession()
	resp, err := session.Request(reqSetting)
	if err != nil {
		return nil, MakeErrorStack(err, "direwolf.Request()")
	}
	return resp, nil
}

// Head is the method to initiate a post request.
// @url: request url, necessary.
func Head(url string, args ...interface{}) (*Response, error) {
	session := NewSession()
	resp, err := session.Head(url, args...)
	if err != nil {
		return nil, MakeErrorStack(err, "direwolf.Post()")
	}
	return resp, nil
}

// Put is the method to initiate a post request.
// @url: request url, necessary.
func Put(url string, args ...interface{}) (*Response, error) {
	session := NewSession()
	resp, err := session.Put(url, args...)
	if err != nil {
		return nil, MakeErrorStack(err, "direwolf.Post()")
	}
	return resp, nil
}

// Patch is the method to initiate a post request.
// @url: request url, necessary.
func Patch(url string, args ...interface{}) (*Response, error) {
	session := NewSession()
	resp, err := session.Patch(url, args...)
	if err != nil {
		return nil, MakeErrorStack(err, "direwolf.Post()")
	}
	return resp, nil
}

// Delete is the method to initiate a post request.
// @url: request url, necessary.
func Delete(url string, args ...interface{}) (*Response, error) {
	session := NewSession()
	resp, err := session.Delete(url, args...)
	if err != nil {
		return nil, MakeErrorStack(err, "direwolf.Post()")
	}
	return resp, nil
}
