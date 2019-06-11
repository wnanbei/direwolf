package direwolf

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/cookiejar"
	"testing"
	"time"

	"golang.org/x/net/publicsuffix"
)

func TestHTTP(t *testing.T) {
	headers := NewHeaders(
		"User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.108 Safari/537.36",
	)
	params := NewParams(
		"keyxxx", "南北",
	)
	cookies := NewCookies(
		"hello", "world",
	)
	// postForm := NewPostForm(
	// 	"this is body", "yes",
	// )
	data := Body("今天天气好")
	// proxy := Proxy("http://127.0.0.1:12333")
	resp, err := Post("https://httpbin.org/post", headers, params, cookies, data)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(resp.Text())
	}
}

func Do(trans *http.Transport, client *http.Client, req *http.Request) {
	// proxyURL, err := url.Parse("http://127.0.0.1:12333")
	// if err != nil {
	// 	panic("proxy url has problem")
	// }
	// trans.Proxy = http.ProxyURL(proxyURL)

	resp, err := client.Do(req)

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(content))
	resp.Body.Close()
}

func TestProxy(t *testing.T) {
	req, err := http.NewRequest("GET", "https://httpbin.org/cookies/set/test/result", nil)
	if err != nil {
		t.Log("fail")
	}
	req.Header.Add("xxxxx", "yyyyy")
	req.AddCookie(&http.Cookie{
		Name:  "hello",
		Value: "world",
	})

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
		panic(err)
	}

	client := &http.Client{
		Transport: trans,
		Jar:       jar,
	}

	Do(trans, client, req)

	req, err = http.NewRequest("GET", "https://httpbin.org/cookies", nil)
	if err != nil {
		t.Log("fail")
	}
	// req.AddCookie(&http.Cookie{
	// 	Name:  "hello",
	// 	Value: "world",
	// })
	Do(trans, client, req)
}
