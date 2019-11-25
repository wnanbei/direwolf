/*
Package direwolf is a convenient and easy to use http client written in Golang.
*/
package direwolf

// Default global session
var session *Session

func init() {
	sessionOptions := DefaultSessionOptions()  // New default global session
	sessionOptions.DisableCookieJar = true
	session = NewSession(sessionOptions)
}

// Send is different with Get and Post method, you should pass a
// Request to it. You can construct Request by use NewRequest
// method.
func Send(req *Request) (*Response, error) {
	resp, err := session.Send(req)
	if err != nil {
		return nil, err
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
// 	direwolf.PostForm: Post data form to send.
// 	direwolf.Body: Post body to send.
// 	direwolf.Proxy: Proxy url to use.
// 	direwolf.Timeout: Request Timeout. Default value is 30.
// 	direwolf.RedirectNum: Number of Request allowed to redirect. Default value is 5.
func Get(URL string, args ...RequestOption) (*Response, error) {
	req := NewRequest("GET", URL, args...)
	resp, err := Send(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Post is the method to constructs and sends a Post request. Parameters are
// the same with direwolf.Get()
//
// Note: direwolf.Body can`t existed with direwolf.PostForm.
func Post(URL string, args ...RequestOption) (*Response, error) {
	req := NewRequest("POST", URL, args...)
	resp, err := Send(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Head is the method to constructs and sends a Head request. Parameters are
// the same with direwolf.Get()
func Head(URL string, args ...RequestOption) (*Response, error) {
	req := NewRequest("HEAD", URL, args...)
	resp, err := Send(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Put is the method to constructs and sends a Put request. Parameters are
// the same with direwolf.Get()
func Put(URL string, args ...RequestOption) (*Response, error) {
	req := NewRequest("Put", URL, args...)
	resp, err := Send(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Patch is the method to constructs and sends a Patch request. Parameters are
// the same with direwolf.Get()
func Patch(URL string, args ...RequestOption) (*Response, error) {
	req := NewRequest("Patch", URL, args...)
	resp, err := Send(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Delete is the method to constructs and sends a Delete request. Parameters are
// the same with direwolf.Get()
func Delete(URL string, args ...RequestOption) (*Response, error) {
	req := NewRequest("Delete", URL, args...)
	resp, err := Send(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
