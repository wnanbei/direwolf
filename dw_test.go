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
	headers := Headers{
		"User-Agent": {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.108 Safari/537.36"},
	}
	params := NewParams(
		"keyxxx", "valuexxx",
	)
	cookies := NewCookies(
		"hello", "world",
	)
	var proxy Proxy = "http://127.0.0.1:12333"
	resp := Get("https://httpbin.org/get", headers, params, cookies, proxy)
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
