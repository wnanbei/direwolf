/*
Package direwolf is a convient and esay to use http client written in Golang.
*/
package direwolf

// Request is different with Get and Post method, you should pass a
// RequestSetting to it. You can construct RequestSetting by use NewRequestSetting
// method.
func Request(reqSetting *RequestSetting) (*Response, error) {
	session := NewSession()
	session.transport.DisableKeepAlives = true
	resp, err := session.Request(reqSetting)
	if err != nil {
		return nil, MakeErrorStack(err, "direwolf.Request()")
	}
	return resp, nil
}

// Get is the most common method of direwolf to constructs and sends a
// Get request.
//
// You can construct this request by passing the following parameters:
// 	url: URL for the request.
// 	http.Header: HTTP Headers to send.
// 	direwolf.Params: Parameters to send in the query string.
// 	direwolf.Cookies: Cookies to send.
// 	direwolf.Postform: Post dataform to send.
// 	direwolf.Body: Post body to send.
// 	direwolf.Proxy: Proxy url to use.
// 	direwolf.Timeout: Request Timeout. Default value is 30.
// 	direwolf.RedirectNum: Number of Request allowed to redirect. Default value is 5.
func Get(URL string, args ...interface{}) (*Response, error) {
	reqSetting := NewRequestSetting("GET", URL, args...)
	resp, err := Request(reqSetting)
	if err != nil {
		return nil, MakeErrorStack(err, "direwolf.Get()")
	}
	return resp, nil
}

// Post is the method to constructs and sends a Post request. Parameters are
// the same with direwolf.Get()
//
// Note: direwolf.Body can`t existed with direwolf.Postform.
func Post(URL string, args ...interface{}) (*Response, error) {
	reqSetting := NewRequestSetting("POST", URL, args...)
	resp, err := Request(reqSetting)
	if err != nil {
		return nil, MakeErrorStack(err, "direwolf.Post()")
	}
	return resp, nil
}

// Head is the method to constructs and sends a Head request. Parameters are
// the same with direwolf.Get()
func Head(URL string, args ...interface{}) (*Response, error) {
	reqSetting := NewRequestSetting("HEAD", URL, args...)
	resp, err := Request(reqSetting)
	if err != nil {
		return nil, MakeErrorStack(err, "direwolf.Post()")
	}
	return resp, nil
}

// Put is the method to constructs and sends a Put request. Parameters are
// the same with direwolf.Get()
func Put(URL string, args ...interface{}) (*Response, error) {
	reqSetting := NewRequestSetting("Put", URL, args...)
	resp, err := Request(reqSetting)
	if err != nil {
		return nil, MakeErrorStack(err, "direwolf.Post()")
	}
	return resp, nil
}

// Patch is the method to constructs and sends a Patch request. Parameters are
// the same with direwolf.Get()
func Patch(URL string, args ...interface{}) (*Response, error) {
	reqSetting := NewRequestSetting("Patch", URL, args...)
	resp, err := Request(reqSetting)
	if err != nil {
		return nil, MakeErrorStack(err, "direwolf.Post()")
	}
	return resp, nil
}

// Delete is the method to constructs and sends a Delete request. Parameters are
// the same with direwolf.Get()
func Delete(URL string, args ...interface{}) (*Response, error) {
	reqSetting := NewRequestSetting("Delete", URL, args...)
	resp, err := Request(reqSetting)
	if err != nil {
		return nil, MakeErrorStack(err, "direwolf.Post()")
	}
	return resp, nil
}
