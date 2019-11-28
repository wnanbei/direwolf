package direwolf

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func newTestProxyServer() *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cli := &http.Client{}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Print("io.ReadFull(r.Body, body) ", err.Error())
		}

		reqUrl := r.URL.String()

		req, err := http.NewRequest(r.Method, reqUrl, strings.NewReader(string(body)))
		if err != nil {
			fmt.Print("http.NewRequest ", err.Error())
			return
		}

		resp, err := cli.Do(req)
		if err != nil {
			fmt.Print("cli.Do(req) ", err.Error())
			return
		}
		defer resp.Body.Close()
		io.Copy(w, resp.Body)
		if _, err := w.Write([]byte("This is proxy Server.")); err != nil {
		}
	}))
	return ts
}

func TestSetProxy(t *testing.T) {
	proxyServer := newTestProxyServer()
	defer proxyServer.Close()
	targetServer := newTestSessionServer()
	defer targetServer.Close()

	proxy := &Proxy{
		HTTP:  proxyServer.URL,
		HTTPS: proxyServer.URL,
	}
	resp, err := Get(targetServer.URL+"/proxy", proxy)
	if err != nil {
		t.Fatal("SetProxy failed: ", err)
	}
	if resp.Text() != "This is target website.This is proxy Server." {
		t.Fatal("SetProxy failed: ", err)
	}
}

func newTestTimeoutServer() *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Second * 2)
		if _, err := w.Write([]byte("timeout")); err != nil {
		}
	}))
	return ts
}

func TestTimeout(t *testing.T) {
	timeoutServer := newTestTimeoutServer()
	defer timeoutServer.Close()

	_, err := Get(timeoutServer.URL, Timeout(1))
	if err != nil {
		fmt.Println(err)
	}
}
