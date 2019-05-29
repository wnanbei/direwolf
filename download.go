package direwolf

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Settings is the Request settings
type Settings struct {
}

// Download is low level request method
func Download(reqSetting *Request, client *http.Client, transport *http.Transport) *Response {
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
	if reqSetting.DataForm != nil {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		data := reqSetting.DataForm.Encode()
		req.Body = ioutil.NopCloser(strings.NewReader(data))
		req.ContentLength = int64(len(data))
	}
	// Handle Cookies
	if reqSetting.Cookies != nil {
		for _, cookie := range reqSetting.Cookies {
			req.AddCookie(cookie)
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
func buildResponse(req *Request, resp *http.Response) *Response {
	return &Response{
		URL:        req.URL,
		StatusCode: resp.StatusCode,
		Proto:      resp.Proto,
		body:       resp.Body,
	}
}
