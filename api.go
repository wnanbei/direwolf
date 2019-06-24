/*
Package direwolf is the most convient and esay to use http client in golang.

You can easily send a request like this:

 import (
	 dw "github.com/wnanbei/direwolf"
 )

 func main() {
	 resp, err := dw.Get("https://www.google.com")
	 if err != nil {
		 ...
	 }
	 fmt.Println(resp.Text())
 }

Besides, direwolf provide a convient way to add parameters to request. Such
as Headers, Cookies, Params, etc.

 import (
	 dw "github.com/wnanbei/direwolf"
 )

 func main() {
	 headers := dw.NewHeaders(
		 "User-Agent", "direwolf",
	 )
	 params := dw.NewParams(
		 "name", "wnanbei",
		 "age", "18",
	 )
	 cookies := dw.NewCookies(
		 "sign", "kzhxciuvyqwekhiuxcyvnkjdhiue",
	 )
	 resp, err := dw.Get("https://httpbin.org/get", headers, params, cookies)
	 if err != nil {
		 ...
	 }
	 fmt.Println(resp.Text())
 }

Output:

 {
   "args": {
     "age": "18",
     "name": "wnanbei"
   },
   "headers": {
     "Accept-Encoding": "gzip",
     "Cookie": "sign=kzhxciuvyqwekhiuxcyvnkjdhiue",
     "Host": "httpbin.org",
     "User-Agent": "direwolf"
   },
   "origin": "118.116.15.151, 118.116.15.151",
   "url": "https://httpbin.org/get?age=18&name=wnanbei"
 }
*/
package direwolf

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
func Get(url string, args ...interface{}) (*Response, error) {
	session := NewSession()
	resp, err := session.Get(url, args...)
	if err != nil {
		return nil, MakeErrorStack(err, "direwolf.Get()")
	}
	return resp, nil
}

// Post is the method to constructs and sends a Post request. Parameters are
// the same with direwolf.Get()
//
// Note: direwolf.Body can`t existed with direwolf.Postform.
func Post(url string, args ...interface{}) (*Response, error) {
	session := NewSession()
	resp, err := session.Post(url, args...)
	if err != nil {
		return nil, MakeErrorStack(err, "direwolf.Post()")
	}
	return resp, nil
}

// Request is different with Get and Post method, you should pass a
// RequestSetting to it. You can construct RequestSetting by use NewRequestSetting
// method.
func Request(reqSetting *RequestSetting) (*Response, error) {
	session := NewSession()
	resp, err := session.Request(reqSetting)
	if err != nil {
		return nil, MakeErrorStack(err, "direwolf.Request()")
	}
	return resp, nil
}

// Head is the method to constructs and sends a Head request. Parameters are
// the same with direwolf.Get()
func Head(url string, args ...interface{}) (*Response, error) {
	session := NewSession()
	resp, err := session.Head(url, args...)
	if err != nil {
		return nil, MakeErrorStack(err, "direwolf.Post()")
	}
	return resp, nil
}

// Put is the method to constructs and sends a Put request. Parameters are
// the same with direwolf.Get()
func Put(url string, args ...interface{}) (*Response, error) {
	session := NewSession()
	resp, err := session.Put(url, args...)
	if err != nil {
		return nil, MakeErrorStack(err, "direwolf.Post()")
	}
	return resp, nil
}

// Patch is the method to constructs and sends a Patch request. Parameters are
// the same with direwolf.Get()
func Patch(url string, args ...interface{}) (*Response, error) {
	session := NewSession()
	resp, err := session.Patch(url, args...)
	if err != nil {
		return nil, MakeErrorStack(err, "direwolf.Post()")
	}
	return resp, nil
}

// Delete is the method to constructs and sends a Delete request. Parameters are
// the same with direwolf.Get()
func Delete(url string, args ...interface{}) (*Response, error) {
	session := NewSession()
	resp, err := session.Delete(url, args...)
	if err != nil {
		return nil, MakeErrorStack(err, "direwolf.Post()")
	}
	return resp, nil
}
