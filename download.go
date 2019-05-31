package direwolf

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Download is low level request method
func Download(reqSetting *RequestSetting, client *http.Client, transport *http.Transport) *Response {
	// New Request
	req, err := http.NewRequest(reqSetting.Method, reqSetting.URL, nil)
	if err != nil {
		panic(err)
	}

	// Add proxy method to transport
	if reqSetting.Proxy != "" {
		proxyURL, err := url.Parse(reqSetting.Proxy)
		if err != nil {
			panic("proxy url has problem")
		}
		transport.Proxy = http.ProxyURL(proxyURL)
	}

	// Handle the Headers.
	req.Header = reqSetting.Headers
	// Handle the DataForm, convert DataForm to strings.Reader.
	// add two new headers: Content-Type and ContentLength.
	if reqSetting.PostForm.data != nil {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		data := reqSetting.PostForm.URLEncode()
		req.Body = ioutil.NopCloser(strings.NewReader(data))
		req.ContentLength = int64(len(data))
	}
	// Handle Cookies
	if reqSetting.Cookies.data != nil {
		for key, values := range reqSetting.Cookies.data {
			for _, value := range values {
				req.AddCookie(&http.Cookie{Name: key, Value: value})
			}
		}
	}

	resp, err := client.Do(req) // do request
	if err != nil {
		panic(err)
	}

	buildedResponse := buildResponse(reqSetting, resp)
	return buildedResponse
}

// buildResponse build response with http.Response after do request.
func buildResponse(req *RequestSetting, resp *http.Response) *Response {
	return &Response{
		URL:        req.URL,
		StatusCode: resp.StatusCode,
		Proto:      resp.Proto,
		body:       resp.Body,
	}
}
