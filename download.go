package direwolf

import (
	"bytes"
	"context"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// send is low level request method.
func send(session *Session, req *Request) (*Response, error) {
	// Set timeout to request context.
	// Default timeout is 30s.
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

	// set RedirectNum to request context.
	// default RedirectNum is 10.
	if req.RedirectNum == 0 {
		ctx = context.WithValue(ctx, "redirectNum", 10)
	} else {
		ctx = context.WithValue(ctx, "redirectNum", req.RedirectNum)
	}

	// Make new http.Request with context
	httpReq, err := http.NewRequestWithContext(ctx, req.Method, req.URL, nil)
	if err != nil {
		return nil, WrapErr(err, "build Request error, please check request url or request method")
	}

	// Handle the Headers.
	httpReq.Header = mergeHeaders(req.Headers, session.Headers)

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
