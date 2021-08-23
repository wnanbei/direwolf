package direwolf

import (
	"net/http"
	"strings"
)

// Request is a prepared request setting, you should construct it by using
// NewRequest().
type Request struct {
	Method        string
	URL           string
	Headers       http.Header
	Body          []byte
	JsonBody      []byte
	Params        *Params
	PostForm      *PostForm
	Cookies       Cookies
	Proxy         *Proxy
	RedirectNum   int
	Timeout       int
	MultipartForm *MultipartForm
}

// NewRequest construct a Request by passing the parameters.
//
// You can construct this request by passing the following parameters:
// 	method: Method for the request.
// 	url: URL for the request.
// 	direwolf.Header: HTTP Headers to send.
// 	direwolf.Params: Parameters to send in the query string.
// 	direwolf.Cookies: Cookies to send.
// 	direwolf.PostForm: Post data form to send.
// 	direwolf.Body: Post body to send.
// 	direwolf.Proxy: Proxy url to use.
// 	direwolf.Timeout: Request Timeout.
// 	direwolf.RedirectNum: Number of Request allowed to redirect.
func NewRequest(method string, URL string, args ...RequestOption) (req *Request, err error) {
	req = &Request{}                     // new a Request and set default field
	req.Method = strings.ToUpper(method) // Upper the method string
	req.URL = URL

	// Check the type of the parameter and handle it.
	for _, arg := range args {
		if err := arg.bindRequest(req); err != nil {
			return nil, WrapErr(err, "set request parameters failed.")
		}
	}
	return req, nil
}
