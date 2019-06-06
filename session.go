package direwolf

import (
	"net"
	"net/http"
	"net/http/cookiejar"
	"time"

	"golang.org/x/net/publicsuffix"
)

// Session is the main object in direwolf. This is its main features:
// 1. handling redirects
// 2. automatically managing cookies
type Session struct {
	client      *http.Client
	transport   *http.Transport
	Headers     http.Header
	Proxy       string
	Timeout     int
	RedirectNum int
}

// Request is a generic request method.
func (session *Session) Request(reqSetting *RequestSetting) (*Response, error) {
	resp, err := session.send(reqSetting)
	if err != nil {
		return nil, MakeErrorStack(err, "direwolf.Session.Request()")
	}
	return resp, nil
}

// Get is a get method.
func (session *Session) Get(URL string, args ...interface{}) (*Response, error) {
	reqSetting := NewRequestSetting("GET", URL, args...)
	resp, err := session.Request(reqSetting)
	if err != nil {
		return nil, MakeErrorStack(err, "direwolf.Session.Request()")
	}
	return resp, nil
}

// Post is a post method.
func (session *Session) Post(URL string, args ...interface{}) (*Response, error) {
	reqSetting := NewRequestSetting("POST", URL, args...)
	resp, err := session.Request(reqSetting)
	if err != nil {
		return nil, MakeErrorStack(err, "direwolf.Session.Request()")
	}
	return resp, nil
}

// NewSession make a Session, and set a default Client and Transport.
func NewSession() *Session {
	trans := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
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
		Transport: trans,
		Jar:       jar,
	}

	return &Session{
		client:    client,
		transport: trans,
	}
}
