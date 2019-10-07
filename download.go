package direwolf

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// send is low level request method.
func (session *Session) send(reqSetting *RequestSetting) (*Response, error) {
	// Make new http.Request
	req, err := http.NewRequest(reqSetting.Method, reqSetting.URL, nil)
	if err != nil {
		return nil, WrapErr(err, "build Request error, please check request url or request method")
	}

	// Handle the Headers.
	req.Header = mergeHeaders(reqSetting.Headers, session.Headers)

	// Add proxy method to transport
	proxyFunc, err := getProxyFunc(reqSetting.Proxy, session.Proxy)
	if err != nil {
		return nil, WrapErr(err, "build proxy failed, please check Proxy and session.Proxy")
	}
	if proxyFunc != nil {
		session.transport.Proxy = proxyFunc
	} else {
		session.transport.Proxy = http.ProxyFromEnvironment
	}

	// set redirect
	session.client.CheckRedirect = getRedirectFunc(reqSetting.RedirectNum)

	// set timeout
	// if timeout > 0, it means a time limit for requests.
	// if timeout < 0, it means no limit.
	// if timeout = 0, it means keep default 30 second timeout.
	if reqSetting.Timeout > 0 {
		session.client.Timeout = time.Duration(reqSetting.Timeout) * time.Second
	} else if reqSetting.Timeout < 0 {
		session.client.Timeout = 0
	} else if session.Timeout > 0 {
		session.client.Timeout = time.Duration(session.Timeout) * time.Second
	} else if session.Timeout < 0 {
		session.client.Timeout = 0
	} else {
		session.client.Timeout = 30 * time.Second
	}

	// Handle the DataForm, convert DataForm to strings.Reader.
	// Set Content-Type to application/x-www-form-urlencoded.
	if reqSetting.Body != nil && reqSetting.PostForm != nil {
		return nil, ErrRequestBody
	} else if reqSetting.PostForm != nil {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		data := reqSetting.PostForm.URLEncode()
		req.Body = ioutil.NopCloser(strings.NewReader(data))
	} else if reqSetting.Body != nil {
		req.Body = ioutil.NopCloser(bytes.NewReader(reqSetting.Body))
	}

	// Handle Cookies
	if reqSetting.Cookies != nil {
		for _, cookie := range reqSetting.Cookies {
			req.AddCookie(cookie)
		}
	}

	resp, err := session.client.Do(req) // do request
	if err != nil {
		return nil, WrapErr(err, "Request Error")
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			panic(err)
		}
	}()

	response, err := buildResponse(reqSetting, resp)
	if err != nil {
		return nil, WrapErr(err, "build Response Error")
	}
	return response, nil
}

// buildResponse build response with http.Response after do request.
func buildResponse(req *RequestSetting, resp *http.Response) (*Response, error) {
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, WrapErr(err, "read Response.Body failed")
	}
	return &Response{
		URL:           req.URL,
		StatusCode:    resp.StatusCode,
		Proto:         resp.Proto,
		Encoding:      "UTF-8",
		Headers:       resp.Header,
		Cookies:       resp.Cookies(),
		Request:       req,
		ContentLength: resp.ContentLength,
		content:       content,
	}, nil
}

// mergeHeaders merge RequestSetting headers and Session Headers.
// RequestSetting has higher priority.
func mergeHeaders(h1, h2 http.Header) http.Header {
	h := http.Header{}
	for key, values := range h2 {
		for _, value := range values {
			h.Set(key, value)
		}
	}
	for key, values := range h1 {
		for _, value := range values {
			h.Set(key, value)
		}
	}
	return h
}

// getProxyFunc return a Proxy Function. RequestSetting has higher priority.
func getProxyFunc(p1, p2 *Proxy) (func(*http.Request) (*url.URL, error), error) {
	var p *Proxy // choose which Proxy to use
	if p1 != nil {
		p = p1
	} else if p2 != nil {
		p = p2
	} else {
		return nil, nil
	}

	if p.HTTP == "" && p.HTTPS == "" { // Check Proxy fields
		return nil, nil
	}

	httpURL, err := url.Parse(p.HTTP)
	if err != nil {
		return nil, WrapErr(err, "Proxy URL error, please check proxy url")
	}
	httpsURL, err := url.Parse(p.HTTPS)
	if err != nil {
		return nil, WrapErr(err, "Proxy URL error, please check proxy url")
	}

	return func(req *http.Request) (*url.URL, error) { // Create a function to choose proxy when transport start request
		if req.URL.Scheme == "http" {
			return httpURL, nil
		} else if req.URL.Scheme == "https" {
			return httpsURL, nil
		}
		err := fmt.Errorf(`unsupported protocol scheme "%s"`, req.URL.Scheme)
		return nil, WrapErr(err, "ProtocolError")
	}, nil
}

// getRedirectFunc return a redirect control function. Default redirect number is 5.
func getRedirectFunc(r int) func(req *http.Request, via []*http.Request) error {
	redirectFunc := func(req *http.Request, via []*http.Request) error {
		if len(via) > r {
			err := &RedirectError{r}
			return WrapErr(err, "RedirectError")
		}
		return nil
	}
	return redirectFunc
}
