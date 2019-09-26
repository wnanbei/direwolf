package direwolf

import (
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"

	"golang.org/x/net/publicsuffix"
)

// Session is the main object in direwolf. This is its main features:
// 1. handling redirects
// 2. automatically managing cookies
type Session struct {
	client    *http.Client
	transport *http.Transport
	Headers   http.Header
	Proxy     *Proxy
	Timeout   int
}

// Request is a generic request method.
func (session *Session) Request(reqSetting *RequestSetting) (*Response, error) {
	resp, err := session.send(reqSetting)
	if err != nil {
		return nil, WrapError(err, "request failed")
	}
	return resp, nil
}

// Get is a get method.
func (session *Session) Get(URL string, args ...interface{}) (*Response, error) {
	reqSetting := NewRequestSetting("GET", URL, args...)
	resp, err := session.Request(reqSetting)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Post is a post method.
func (session *Session) Post(URL string, args ...interface{}) (*Response, error) {
	reqSetting := NewRequestSetting("POST", URL, args...)
	resp, err := session.Request(reqSetting)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Head is a post method.
func (session *Session) Head(URL string, args ...interface{}) (*Response, error) {
	reqSetting := NewRequestSetting("HEAD", URL, args...)
	resp, err := session.Request(reqSetting)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Put is a post method.
func (session *Session) Put(URL string, args ...interface{}) (*Response, error) {
	reqSetting := NewRequestSetting("PUT", URL, args...)
	resp, err := session.Request(reqSetting)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Patch is a post method.
func (session *Session) Patch(URL string, args ...interface{}) (*Response, error) {
	reqSetting := NewRequestSetting("PATCH", URL, args...)
	resp, err := session.Request(reqSetting)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Delete is a post method.
func (session *Session) Delete(URL string, args ...interface{}) (*Response, error) {
	reqSetting := NewRequestSetting("DELETE", URL, args...)
	resp, err := session.Request(reqSetting)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Cookies returns the cookies of the given url in Session.
func (session *Session) Cookies(URL string) Cookies {
	if session.client.Jar == nil {
		return nil
	}
	parsedURL, err := url.Parse(URL)
	if err != nil {
		return nil
	}
	return session.client.Jar.Cookies(parsedURL)
}

// SetCookies set cookies of the url in Session.
func (session *Session) SetCookies(URL string, cookies Cookies) {
	if session.client.Jar == nil {
		return
	}
	parsedURL, err := url.Parse(URL)
	if err != nil {
		return
	}
	session.client.Jar.SetCookies(parsedURL, cookies)
}

// DisableCookieJar disable the CookieJar of session
func (session *Session) DisableCookieJar() {
	session.client.Jar = nil
}

// NewSession new a Session object, and set a default Client and Transport.
func NewSession() *Session {
	defaultTransport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		MaxIdleConnsPerHost:   2,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	// set CookieJar
	options := cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	}
	jar, err := cookiejar.New(&options)
	if err != nil {
		return nil
	}

	client := &http.Client{
		Transport: defaultTransport,
		Jar:       jar,
	}

	headers := http.Header{}
	headers.Add("User-Agent", "direwolf - winter is coming")

	return &Session{
		client:    client,
		transport: defaultTransport,
		Headers:   headers,
	}
}
