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

// Send is a generic request method.
func (session *Session) Send(req *Request) (*Response, error) {
	resp, err := session.send(req)
	if err != nil {
		return nil, WrapErr(err, "request failed")
	}
	return resp, nil
}

// Get is a get method.
func (session *Session) Get(URL string, args ...RequestOption) (*Response, error) {
	req := NewRequest("GET", URL, args...)
	resp, err := session.Send(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Post is a post method.
func (session *Session) Post(URL string, args ...RequestOption) (*Response, error) {
	req := NewRequest("POST", URL, args...)
	resp, err := session.Send(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Head is a post method.
func (session *Session) Head(URL string, args ...RequestOption) (*Response, error) {
	req := NewRequest("HEAD", URL, args...)
	resp, err := session.Send(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Put is a post method.
func (session *Session) Put(URL string, args ...RequestOption) (*Response, error) {
	req := NewRequest("PUT", URL, args...)
	resp, err := session.Send(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Patch is a post method.
func (session *Session) Patch(URL string, args ...RequestOption) (*Response, error) {
	req := NewRequest("PATCH", URL, args...)
	resp, err := session.Send(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Delete is a post method.
func (session *Session) Delete(URL string, args ...RequestOption) (*Response, error) {
	req := NewRequest("DELETE", URL, args...)
	resp, err := session.Send(req)
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

// NewSession new a Session object, and set a default Client and Transport.
func NewSession(options ...*SessionOptions) *Session {
	var sessionOptions *SessionOptions
	if len(options) > 0 {
		sessionOptions = options[0]
	} else {
		sessionOptions = DefaultSessionOptions()
	}

	// set transport parameters.
	trans := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   sessionOptions.DialTimeout,
			KeepAlive: sessionOptions.DialKeepAlive,
		}).DialContext,
		MaxIdleConns:          sessionOptions.MaxIdleConns,
		MaxIdleConnsPerHost:   sessionOptions.MaxIdleConnsPerHost,
		MaxConnsPerHost:       sessionOptions.MaxConnsPerHost,
		IdleConnTimeout:       sessionOptions.IdleConnTimeout,
		TLSHandshakeTimeout:   sessionOptions.TLSHandshakeTimeout,
		ExpectContinueTimeout: sessionOptions.ExpectContinueTimeout,
	}
	if sessionOptions.DisableDialKeepAlives {
		trans.DisableKeepAlives = true
	}

	client := &http.Client{Transport: trans}

	// set CookieJar
	if sessionOptions.DisableCookieJar == false {
		cookieJarOptions := cookiejar.Options{
			PublicSuffixList: publicsuffix.List,
		}
		jar, err := cookiejar.New(&cookieJarOptions)
		if err != nil {
			return nil
		}
		client.Jar = jar
	}

	// Set default user agent
	headers := http.Header{}
	headers.Add("User-Agent", "direwolf - winter is coming")

	return &Session{
		client:    client,
		transport: trans,
		Headers:   headers,
	}
}

type SessionOptions struct {
	// DialTimeout is the maximum amount of time a dial will wait for
	// a connect to complete.
	//
	// When using TCP and dialing a host name with multiple IP
	// addresses, the timeout may be divided between them.
	//
	// With or without a timeout, the operating system may impose
	// its own earlier timeout. For instance, TCP timeouts are
	// often around 3 minutes.
	DialTimeout time.Duration

	// DialKeepAlive specifies the interval between keep-alive
	// probes for an active network connection.
	//
	// Network protocols or operating systems that do
	// not support keep-alives ignore this field.
	// If negative, keep-alive probes are disabled.
	DialKeepAlive time.Duration

	// MaxConnsPerHost optionally limits the total number of
	// connections per host, including connections in the dialing,
	// active, and idle states. On limit violation, dials will block.
	//
	// Zero means no limit.
	MaxConnsPerHost int

	// MaxIdleConns controls the maximum number of idle (keep-alive)
	// connections across all hosts. Zero means no limit.
	MaxIdleConns int

	// MaxIdleConnsPerHost, if non-zero, controls the maximum idle
	// (keep-alive) connections to keep per-host. If zero,
	// DefaultMaxIdleConnsPerHost is used.
	MaxIdleConnsPerHost int

	// IdleConnTimeout is the maximum amount of time an idle
	// (keep-alive) connection will remain idle before closing
	// itself.
	// Zero means no limit.
	IdleConnTimeout time.Duration

	// TLSHandshakeTimeout specifies the maximum amount of time waiting to
	// wait for a TLS handshake. Zero means no timeout.
	TLSHandshakeTimeout time.Duration

	// ExpectContinueTimeout, if non-zero, specifies the amount of
	// time to wait for a server's first response headers after fully
	// writing the request headers if the request has an
	// "Expect: 100-continue" header. Zero means no timeout and
	// causes the body to be sent immediately, without
	// waiting for the server to approve.
	// This time does not include the time to send the request header.
	ExpectContinueTimeout time.Duration

	// DisableCookieJar specifies whether disable session cookiejar.
	DisableCookieJar bool

	// DisableDialKeepAlives, if true, disables HTTP keep-alives and
	// will only use the connection to the server for a single
	// HTTP request.
	//
	// This is unrelated to the similarly named TCP keep-alives.
	DisableDialKeepAlives bool
}

// DefaultSessionOptions return a default SessionOptions object.
func DefaultSessionOptions() *SessionOptions {
	return &SessionOptions{
		DialTimeout:           30 * time.Second,
		DialKeepAlive:         30 * time.Second,
		MaxConnsPerHost:       0,
		MaxIdleConns:          100,
		MaxIdleConnsPerHost:   2,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		DisableCookieJar:      false,
		DisableDialKeepAlives: false,
	}
}
