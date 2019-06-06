package direwolf

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// send is low level request method
func (session *Session) send(reqSetting *RequestSetting) (*Response, error) {
	// Make new http.Request
	req, err := http.NewRequest(reqSetting.Method, reqSetting.URL, nil)
	if err != nil {
		return nil, MakeError(err, "NewRequestError", "Build Request error, please check request url or request method")
	}

	// Add proxy method to transport
	proxyFunc, err := getProxyFunc(reqSetting.Proxy, session.Proxy)
	if err != nil {
		return nil, MakeErrorStack(err, "direwolf.Session.send()")
	}
	if proxyFunc != nil {
		session.Transport.Proxy = proxyFunc
	} else {
		session.Transport.Proxy = http.ProxyFromEnvironment
	}

	// set timeout
	// if timeout > 0, it means a time limit for requests.
	// if timeout < 0, it means no limit.
	// if timeout = 0, it means keep default 30 second timeout.
	if reqSetting.Timeout > 0 {
		session.Client.Timeout = time.Duration(reqSetting.Timeout) * time.Second
	} else if reqSetting.Timeout < 0 {
		session.Client.Timeout = 0
	} else if session.Timeout > 0 {
		session.Client.Timeout = time.Duration(session.Timeout) * time.Second
	} else if session.Timeout < 0 {
		session.Client.Timeout = 0
	} else {
		session.Client.Timeout = 30 * time.Second
	}

	// Handle the Headers.
	req.Header = mergeHeaders(reqSetting.Headers, session.Headers)

	// Handle the DataForm, convert DataForm to strings.Reader.
	// Set Content-Type to application/x-www-form-urlencoded.
	if reqSetting.Body != nil && reqSetting.PostForm != nil {
		return nil, MakeError(nil, "RequestBodyError", "Body can`t exists with PostForm")
	} else if reqSetting.PostForm != nil {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		data := reqSetting.PostForm.URLEncode()
		req.Body = ioutil.NopCloser(strings.NewReader(data))
	} else if reqSetting.Body != nil {
		req.Body = ioutil.NopCloser(bytes.NewReader(reqSetting.Body))
	}

	// Handle Cookies
	if reqSetting.Cookies != nil {
		for key, values := range reqSetting.Cookies.data {
			for _, value := range values {
				req.AddCookie(&http.Cookie{Name: key, Value: value})
			}
		}
	}

	resp, err := session.Client.Do(req) // do request
	if err != nil {
		return nil, MakeError(err, "HTTPError", "Request Error")
	}

	buildedResponse := buildResponse(reqSetting, resp)
	return buildedResponse, nil
}

// buildResponse build response with http.Response after do request.
func buildResponse(req *RequestSetting, resp *http.Response) *Response {
	return &Response{
		URL:        req.URL,
		StatusCode: resp.StatusCode,
		Proto:      resp.Proto,
		body:       resp.Body,
	}
}

func mergeHeaders(h1, h2 http.Header) http.Header {
	h := http.Header{}
	if h1 != nil && h2 != nil {
		for key, values := range h1 {
			for _, value := range values {
				h.Add(key, value)
			}
		}
		for key, values := range h2 {
			for _, value := range values {
				h.Add(key, value)
			}
		}
	} else if h1 != nil {
		h = h1
	} else if h2 != nil {
		h = h2
	}
	return h
}

func getProxyFunc(p1, p2 string) (func(*http.Request) (*url.URL, error), error) {
	// Add proxy method to transport
	var p string
	if p1 != "" {
		p = p1
	} else if p2 != "" {
		p = p2
	} else {
		return nil, nil
	}

	proxyURL, err := url.Parse(p)
	if err != nil {
		return nil, MakeError(err, "ProxyURLError", "Proxy URL error, please check proxy url")
	}
	return http.ProxyURL(proxyURL), nil
}