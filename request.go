package direwolf

import (
	"net/http"
	"strings"
)

// RequestSetting is a prepared request setting, you should construct it by using
// NewRequestSetting().
type RequestSetting struct {
	Method      string
	URL         string
	Headers     http.Header
	Body        Body
	Params      *Params
	PostForm    *PostForm
	Cookies     *Cookies
	Proxy       *Proxy
	RedirectNum int
	Timeout     int
}

// NewRequestSetting construct a RequestSetting by passing the parameters.
//
// You can construct this request by passing the following parameters:
// 	method: Method for the request.
// 	url: URL for the request.
// 	http.Header: HTTP Headers to send.
// 	direwolf.Params: Parameters to send in the query string.
// 	direwolf.Cookies: Cookies to send.
// 	direwolf.Postform: Post dataform to send.
// 	direwolf.Body: Post body to send.
// 	direwolf.Proxy: Proxy url to use.
// 	direwolf.Timeout: Request Timeout.
// 	direwolf.RedirectNum: Number of Request allowed to redirect.
func NewRequestSetting(method string, URL string, args ...interface{}) *RequestSetting {
	reqSetting := &RequestSetting{}             // new a RequestSetting and set default field
	reqSetting.Method = strings.ToUpper(method) // Upper the method string
	reqSetting.URL = URL
	reqSetting.RedirectNum = 5

	// Check the type of the parameter and handle it.
	for _, arg := range args {
		switch a := arg.(type) {
		case http.Header:
			reqSetting.Headers = a
		case *Params:
			reqSetting.Params = a
			reqSetting.URL = reqSetting.URL + "?" + reqSetting.Params.URLEncode()
		case *PostForm:
			reqSetting.PostForm = a
		case Body:
			reqSetting.Body = a
		case *Cookies:
			reqSetting.Cookies = a
		case *Proxy:
			reqSetting.Proxy = a
		case RedirectNum:
			reqSetting.RedirectNum = int(a)
		case Timeout:
			reqSetting.Timeout = int(a)
		}
	}
	return reqSetting
}
