package direwolf

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func newTestProxyServer() *httptest.Server {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.GET("/:a", func(c *gin.Context) {
		reqUrl := c.Request.URL.String() // Get request url
		req, err := http.NewRequest(c.Request.Method, reqUrl, nil)
		if err != nil {
			c.AbortWithStatus(404)
			return
		}

		// Forwarding requests from client.
		cli := &http.Client{}
		resp, err := cli.Do(req)
		if err != nil {
			c.AbortWithStatus(404)
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		c.Data(200, "text/plain", body)        // write response body to response.
		c.String(200, "This is proxy Server.") // add proxy info.
	})
	ts := httptest.NewServer(router)
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
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.GET("/", func(c *gin.Context) {
		time.Sleep(time.Second * 2)
		c.String(200, "successed")
	})
	ts := httptest.NewServer(router)
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
}
