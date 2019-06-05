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
	Client    *http.Client
	Transport *http.Transport
	Cookies   *cookiejar.Jar
	Headers   *http.Header
	Proxy     string
	Timeout   int
}

// Request is a generic request method.
func (session *Session) Request(method string, URL string, args ...interface{}) (*Response, error) {
	preq := NewRequestSetting(method, URL, args...)
	resp, err := session.send(preq)
	if err != nil {
		return nil, MakeErrorStack(err, "direwolf.Session.Request()")
	}
	return resp, nil
}

// Get is a get method.
func (session *Session) Get(URL string, args ...interface{}) (*Response, error) {
	return session.Request("GET", URL, args...)
}

// Post is a post method.
func (session *Session) Post(URL string, args ...interface{}) (*Response, error) {
	return session.Request("POST", URL, args...)
}

// send is responsible for handling some subsequent processing of the PreRequest.
func (session *Session) send(preq *RequestSetting) (*Response, error) {
	response, err := Download(preq, session.Client, session.Transport)
	if err != nil {
		return nil, MakeErrorStack(err, "direwolf.Session.send()")
	}
	// build response
	return response, nil
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
		Client:    client,
		Transport: trans,
	}
}
