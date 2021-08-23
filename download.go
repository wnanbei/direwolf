package direwolf

import (
	"bytes"
	"context"
	"errors"
	"fmt"
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
	} else if req.RedirectNum > 0 {
		ctx = context.WithValue(ctx, "redirectNum", req.RedirectNum)
	} else {
		ctx = context.WithValue(ctx, "redirectNum", 0)
	}

	// Make new http.Request with context
	httpReq, err := http.NewRequestWithContext(ctx, req.Method, req.URL, nil)
	if err != nil {
		timeoutCancel()
		return nil, WrapErr(err, "build Request error, please check request url or request method")
	}

	// Handle the Headers.
	httpReq.Header = mergeHeaders(req.Headers, session.Headers)

	// Handle the DataForm, Body or JsonBody.
	// Set right Content-Type.
	if req.PostForm != nil {
		httpReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		data := req.PostForm.URLEncode()
		httpReq.Body = ioutil.NopCloser(strings.NewReader(data))
	} else if req.Body != nil {
		httpReq.Body = ioutil.NopCloser(bytes.NewReader(req.Body))
	} else if req.JsonBody != nil {
		httpReq.Header.Set("Content-Type", "application/json")
		httpReq.Body = ioutil.NopCloser(bytes.NewReader(req.JsonBody))
	} else if req.MultipartForm != nil {
		ct := fmt.Sprintf("multipart/form-data; boundary=--%s", req.MultipartForm.Boundary())
		httpReq.Header.Set("Content-Type", ct)
		httpReq.Body = ioutil.NopCloser(req.MultipartForm.Reader())
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
			timeoutCancel()
			return nil, WrapErr(ErrTimeout, err.Error())
		}
		timeoutCancel()
		return nil, WrapErr(err, "Request Error")
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			panic(err)
		}
	}()

	response, err := buildResponse(req, resp)
	if err != nil {
		timeoutCancel()
		return nil, WrapErr(err, "build Response Error")
	}

	timeoutCancel() // cancel the timeout context after request succeed.
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
