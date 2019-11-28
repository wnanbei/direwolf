package direwolf

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
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
		if !errors.Is(err, ErrTimeout) {
			t.Fatal("TestTimeout failed: ", err)
		}
	}

	_, err = Get(timeoutServer.URL, Timeout(3))
	if err != nil {
		t.Fatal("TestTimeout failed: ", err)
	}
}

func newTestRedirectServer() *httptest.Server {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.GET("/", func(c *gin.Context) {
		c.String(200, "successed")
	})
	router.GET("/:number", func(c *gin.Context) {
		number, _ := strconv.Atoi(c.Param("number"))
		if number == 1 {
			c.Redirect(302, "/")
		}
		number = number - 1
		c.Redirect(302, strconv.Itoa(number))
	})
	ts := httptest.NewServer(router)
	return ts
}

func TestRedirect(t *testing.T) {
	redirectServer := newTestRedirectServer()
	defer redirectServer.Close()

	redirect := RedirectNum(-1) // ban redirect
	resp, err := Get(redirectServer.URL+"/", redirect)
	if err != nil {
		t.Fatal("Test TestRedirectError failed.")
	}
	if resp.Text() != "successed" {
		t.Fatal("Test TestRedirectError failed.")
	}

	redirect = RedirectNum(1) // allow 1 redirect
	resp, err = Get(redirectServer.URL+"/1", redirect)
	if err != nil {
		t.Fatal("Test TestRedirectError failed.")
	}
	if resp.Text() != "successed" {
		t.Fatal("Test TestRedirectError failed.")
	}

	redirect = RedirectNum(1) // allow 1 redirect
	resp, err = Get(redirectServer.URL+"/2", redirect)
	if err != nil {
		var errType *RedirectError
		if !errors.As(err, &errType) { // check RedirectError
			fmt.Println("yyyyy")
			t.Fatal("Test TestRedirectError failed.")
		}
	} else {
		t.Fatal("Test TestRedirectError failed.")
	}
	if resp.Text() != "successed" {
		t.Fatal("Test TestRedirectError failed.")
	}
}
