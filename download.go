package direwolf

import (
	"bytes"
	"context"
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
func send(session *Session, req *Request) (*Response, error) {
	// set timeout to request context.
	timeout := time.Second * 30
	if req.Timeout > 0 {
		timeout = time.Second * time.Duration(req.Timeout)
	} else if session.Timeout > 0 {
		timeout = time.Second * time.Duration(session.Timeout)
	}
	ctx, timeoutCancel := context.WithTimeout(context.Background(), timeout)

	// set proxy to request context.
	if req.Proxy != nil {
		ctx = context.WithValue(ctx, "http", req.Proxy.HTTP)
		ctx = context.WithValue(ctx, "https", req.Proxy.HTTPS)
	} else if session.Proxy != nil {
		ctx = context.WithValue(ctx, "http", session.Proxy.HTTP)
		ctx = context.WithValue(ctx, "https", session.Proxy.HTTPS)
	}

	// Make new http.Request with context
	httpReq, err := http.NewRequestWithContext(ctx, req.Method, req.URL, nil)
	if err != nil {
		return nil, WrapErr(err, "build Request error, please check request url or request method")
	}

	// Handle the Headers.
	httpReq.Header = mergeHeaders(req.Headers, session.Headers)

	// set redirect
	session.client.CheckRedirect = getRedirectFunc(req.RedirectNum)

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
		if strings.Contains(err.Error(), "context deadline exceeded") { // check timeout error
			return nil, WrapErr(ErrTimeout, err.Error())
		}
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

	timeoutCancel() // cancel the timeout context after request successed.
	return response, nil
}

// buildResponse build response with http.Response after do request.
func buildResponse(httpReq *Request, httpResp *http.Response) (*Response, error) {
	content, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		if !errors.Is(err, io.ErrUnexpectedEOF) { // Ignore Unexpected EOF error
			return nil, WrapErr(err, "read Response.Body failed")
		}
	}
	return &Response{
		URL:           httpReq.URL,
		StatusCode:    httpResp.StatusCode,
		Proto:         httpResp.Proto,
		Headers:       httpResp.Header,
		Cookies:       httpResp.Cookies(),
		Request:       httpReq,
		ContentLength: httpResp.ContentLength,
		Content:       content,
		encoding:      "UTF-8",
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
