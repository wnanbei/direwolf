package direwolf

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// send is low level request method.
func (session *Session) send(req *Request) (*Response, error) {
	// Make new http.Request
	httpReq, err := http.NewRequest(req.Method, req.URL, nil)
	if err != nil {
		return nil, WrapErr(err, "build Request error, please check request url or request method")
	}

	// Handle the Headers.
	httpReq.Header = mergeHeaders(req.Headers, session.Headers)

	// Add proxy method to transport
	proxyFunc, err := getProxyFunc(req.Proxy, session.Proxy)
	if err != nil {
		return nil, WrapErr(err, "build proxy failed, please check Proxy and session.Proxy")
	}
	if proxyFunc != nil {
		session.transport.Proxy = proxyFunc
	} else {
		session.transport.Proxy = http.ProxyFromEnvironment
	}

	// set redirect
	session.client.CheckRedirect = getRedirectFunc(req.RedirectNum)

	// set timeout
	// if timeout > 0, it means a time limit for requests.
	// if timeout < 0, it means no limit.
	// if timeout = 0, it means keep default 30 second timeout.
	if req.Timeout > 0 {
		session.client.Timeout = time.Duration(req.Timeout) * time.Second
	} else if req.Timeout < 0 {
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
	if req.Body != nil && req.PostForm != nil {
		return nil, ErrRequestBody
	} else if req.PostForm != nil {
		httpReq.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		data := req.PostForm.URLEncode()
		httpReq.Body = ioutil.NopCloser(strings.NewReader(data))
	} else if req.Body != nil {
		httpReq.Body = ioutil.NopCloser(bytes.NewReader(req.Body))
	}

	// Handle Cookies
	if req.Cookies != nil {
		for _, cookie := range req.Cookies {
			httpReq.AddCookie(cookie)
		}
	}

	resp, err := session.client.Do(httpReq) // do request
	if err != nil {
		return nil, WrapErr(err, "Request Error")
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			panic(err)
		}
	}()

	response, err := buildResponse(req, resp)
	if err != nil {
		return nil, WrapErr(err, "build Response Error")
	}
	return response, nil
}

// buildResponse build response with http.Response after do request.
func buildResponse(httpReq *Request, httpResp *http.Response) (*Response, error) {
	content, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		if !errors.Is(err, io.ErrUnexpectedEOF) {  // Ignore Unexpected EOF error
			return nil, WrapErr(err, "read Response.Body failed")
		}
	}
	return &Response{
		URL:           httpReq.URL,
		StatusCode:    httpResp.StatusCode,
		Proto:         httpResp.Proto,
		Encoding:      "UTF-8",
		Headers:       httpResp.Header,
		Cookies:       httpResp.Cookies(),
		Request:       httpReq,
		ContentLength: httpResp.ContentLength,
		content:       content,
	}, nil
}

// mergeHeaders merge Request headers and Session Headers.
// Request has higher priority.
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

// getProxyFunc return a Proxy Function. Request has higher priority.
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
