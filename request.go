package direwolf

import (
	"net/http"
	"net/url"
	"fmt"
	"io/ioutil"
	"strings"
)

type PreRequest struct{
	Method string
	Url string
	Headers Headers
	Cookies Cookies
	Data Data
	DataForm url.Values
	Params url.Values
}

type Session struct{
	client  *http.Client
}
func (self Session) prepareRequest(method string, reqUrl string, args ...interface{}) *PreRequest {
	req := new(PreRequest)
	req.Method = method
	req.Url = reqUrl
	for _, arg := range(args){
		switch a := arg.(type) {
		case Headers:
			req.Headers = make(map[string]string)
			for key, value := range a {
				req.Headers[key] = value
			}
		case Params:
			req.Params = url.Values{}
			for key, value := range a {
				req.Params.Add(key, value)
			}
		case Cookies:
			req.Cookies = make(map[string]string)
			for key, value := range a {
				req.Cookies[key] = value
			}
		case DataForm:
			req.DataForm = url.Values{}
			for key, value := range a {
				req.DataForm.Add(key, value)
			}
		case Data:
			req.Data = a
		}
	}
	return req
}
func (self *Session) Request(method string, reqUrl string, args ...interface{}) {
	preq := self.prepareRequest(method, reqUrl, args...)
	self.send(preq)
}
func (self *Session) Get(reqUrl string, args ...interface{}){
	self.Request("Get", reqUrl, args...)
}
func (self *Session) Post(reqUrl string, args ...interface{}){
	self.Request("Post", reqUrl, args...)
}
func (self *Session) send(preq *PreRequest) *Response {
	self.client = &http.Client{}
	req, err := http.NewRequest(preq.Method, preq.Url, nil)
	if err != nil {
		panic(err)
	}

	for key, value := range preq.Headers{
		req.Header.Add(key, value)
	}
	if preq.DataForm != nil {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		data := preq.DataForm.Encode()
		req.Body = ioutil.NopCloser(strings.NewReader(data))
		req.ContentLength = int64(len(data))
	}

	resp, err := self.client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	return &Response{}
}