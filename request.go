package direwolf

import (
	"net/http"
	"strings"
)

// Request is a prepared request setting, you should construct it by using
// NewRequest().
type Request struct {
	Method      string
	URL         string
	Headers     http.Header
	Body        Body
	Params      *Params
	PostForm    *PostForm
	Cookies     Cookies
	Proxy       *Proxy
	RedirectNum int
	Timeout     int
}

// NewRequest construct a Request by passing the parameters.
//
// You can construct this request by passing the following parameters:
// 	method: Method for the request.
// 	url: URL for the request.
// 	http.Header: HTTP Headers to send.
// 	direwolf.Params: Parameters to send in the query string.
// 	direwolf.Cookies: Cookies to send.
// 	direwolf.PostForm: Post data form to send.
// 	direwolf.Body: Post body to send.
// 	direwolf.Proxy: Proxy url to use.
// 	direwolf.Timeout: Request Timeout.
// 	direwolf.RedirectNum: Number of Request allowed to redirect.
func NewRequest(method string, URL string, args ...interface{}) *Request {
	req := &Request{}             // new a Request and set default field
	req.Method = strings.ToUpper(method) // Upper the method string
	req.URL = URL
	req.RedirectNum = 5

	// Check the type of the parameter and handle it.
	for _, arg := range args {
		switch a := arg.(type) {
		case http.Header:
			req.Headers = a
		case *Params:
			req.Params = a
			req.URL = req.URL + "?" + req.Params.URLEncode()
		case *PostForm:
			req.PostForm = a
		case Body:
			req.Body = a
		case Cookies:
			req.Cookies = a
		case *Proxy:
			req.Proxy = a
		case RedirectNum:
			req.RedirectNum = int(a)
		case Timeout:
			req.Timeout = int(a)
		}
	}
	return req
}
