package direwolf

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"testing"
	"time"
)

func TestHTTP(t *testing.T) {
	headers := NewHeaders(
		"User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.108 Safari/537.36",
	)
	params := NewParams(
		"keyxxx", "valuexxx",
	)
	cookies := NewCookies(
		"hello", "world",
	)
	data := NewPostForm(
		"this is body", "yes",
	)
	proxy := Proxy("http://127.0.0.1:12333")
	resp := Post("https://httpbin.org/post", headers, params, cookies, data, proxy)
	t.Log(resp.Text())
}

func Do(trans *http.Transport, client *http.Client, req *http.Request) {
	proxyURL, err := url.Parse("http://127.0.0.1:12333")
	if err != nil {
		panic("proxy url has problem")
	}
	trans.Proxy = http.ProxyURL(proxyURL)

	resp, err := client.Do(req)

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(content))
	resp.Body.Close()
}

func TestProxy(t *testing.T) {
	req, err := http.NewRequest("GET", "https://httpbin.org/get", nil)
	if err != nil {
		t.Log("fail")
	}
	req.Header.Add("xxxxx", "yyyyy")

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
	client := &http.Client{
		Transport: trans,
	}

	Do(trans, client, req)
}

// type headers http.Header

// func (h *headers) New(keyvalue ...string) {
// 	h = http.Header{}
// 	if keyvalue != nil {
// 		if len(keyvalue)%2 != 0 {
// 			panic("key and value must be part")
// 		}

// 		for i := 0; i < len(keyvalue)/2; i++ {
// 			key := keyvalue[i*2]
// 			value := keyvalue[i*2+1]
// 			h[key] = append(h.data[key], value)
// 		}
// 	}
// }

// func TestHeader(t *testing.T) {
// 	h := &headers{}
// 	h.New("dxxx", "eeee", "wwww", "xxxx")
// 	h.Add("a", "b")
// 	t.Log(h)
// }
