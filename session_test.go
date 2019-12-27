package direwolf

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func newTestSessionServer() *httptest.Server {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.GET("/test", func(c *gin.Context) {
		c.String(200, "GET")
	})
	router.GET("/setCookie", func(c *gin.Context) {
		http.SetCookie(c.Writer, &http.Cookie{
			Name:    "key",
			Value:   "value",
			Path:    "/",
			Expires: time.Now().Add(30 * time.Second),
		})
	})
	router.GET("/getCookie", func(c *gin.Context) {
		cookies := c.Request.Cookies()
		for _, cookie := range cookies {
			c.String(200, cookie.Name+"="+cookie.Value)
		}
	})
	router.GET("/getHeader", func(c *gin.Context) {
		header := c.GetHeader("user-agent")
		c.String(200, header)
	})
	router.GET("/getParams", func(c *gin.Context) {
		value := c.Query("key")
		c.String(200, value)
	})
	router.GET("/proxy", func(c *gin.Context) {
		c.String(200, "This is target website.")
	})
	router.POST("/test", func(c *gin.Context) {
		data, err := c.GetRawData()
		if err != nil {
			c.AbortWithStatus(404)
			return
		}
		if string(data) == "key=value" {
			c.String(200, "POST")
		} else {
			c.String(200, "Failed")
		}
	})
	router.HEAD("/", func(c *gin.Context) {
		c.SetCookie("HEAD", "RIGHT", 1000, "/", "localhost", false, false)
	})
	router.PUT("/", func(c *gin.Context) {
		c.String(200, "PUT")
	})
	router.DELETE("/", func(c *gin.Context) {
		c.String(200, "DELETE")
	})
	router.PATCH("/", func(c *gin.Context) {
		c.String(200, "PATCH")
	})

	ts := httptest.NewServer(router)
	return ts

	//ys := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//	// check method is GET before going to check other features
	//	if r.Method == "GET" {
	//		if r.URL.Path == "/test" {
	//			if _, err := w.Write([]byte("GET")); err != nil {
	//			}
	//		}
	//		if r.URL.Path == "/setCookie" {
	//			http.SetCookie(w, &http.Cookie{Name: "key", Value: "value"})
	//		}
	//		if r.URL.Path == "/getCookie" {
	//			cookies := r.Cookies()
	//			for _, cookie := range cookies {
	//				if _, err := w.Write([]byte(cookie.Name + "=" + cookie.Value)); err != nil {
	//				}
	//			}
	//		}
	//		if r.URL.Path == "/getHeader" {
	//			header := r.Header
	//			value := header.Get("user-agent")
	//			if _, err := w.Write([]byte(value)); err != nil {
	//			}
	//		}
	//		if r.URL.Path == "/getParams" {
	//			if err := r.ParseForm(); err != nil {
	//			}
	//			params := r.Form
	//			value := params.Get("key")
	//			if _, err := w.Write([]byte(value)); err != nil {
	//			}
	//		}
	//		if r.URL.Path == "/proxy" {
	//			if _, err := w.Write([]byte("This is target website.")); err != nil {
	//			}
	//		}
	//	}
	//	if r.Method == "POST" {
	//		if r.URL.Path == "/test" {
	//			body, _ := ioutil.ReadAll(r.Body)
	//			if string(body) != "key=value" {
	//				if _, err := w.Write([]byte("Failed")); err != nil {
	//				}
	//			} else {
	//				if _, err := w.Write([]byte("POST")); err != nil {
	//				}
	//			}
	//		}
	//	}
	//	if r.Method == "HEAD" {
	//		http.SetCookie(w, &http.Cookie{Name: "HEAD", Value: "RIGHT"})
	//	}
	//	if r.Method == "PUT" {
	//		if _, err := w.Write([]byte("PUT")); err != nil {
	//		}
	//	}
	//	if r.Method == "PATCH" {
	//		if _, err := w.Write([]byte("PATCH")); err != nil {
	//		}
	//	}
	//	if r.Method == "DELETE" {
	//		if _, err := w.Write([]byte("DELETE")); err != nil {
	//		}
	//	}
	//}))
	//return ys
}

func TestSessionGet(t *testing.T) {
	ts := newTestSessionServer()
	defer ts.Close()

	session := NewSession()
	resp, err := session.Get(ts.URL + "/test")
	if err != nil {
		t.Fatal(err)
	}
	text := resp.Text()
	if text != "GET" {
		t.Fatal("Session.Get test failed")
	}
}

func TestSessionPost(t *testing.T) {
	ts := newTestSessionServer()
	defer ts.Close()

	session := NewSession()
	postForm := NewPostForm(
		"key", "value",
	)
	resp, err := session.Post(ts.URL+"/test", postForm)
	if err != nil {
		t.Fatal(err)
	}
	text := resp.Text()
	if text != "POST" {
		t.Fatal("Session.Post test failed")
	}

	body := Body("key=value")
	resp2, err := session.Post(ts.URL+"/test", body)
	if err != nil {
		t.Fatal(err)
	}
	text2 := resp2.Text()
	if text2 != "POST" {
		t.Fatal("Session.Post test failed")
	}
}

func TestSessionPut(t *testing.T) {
	ts := newTestSessionServer()
	defer ts.Close()

	session := NewSession()
	resp, err := session.Put(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	text := resp.Text()
	if text != "PUT" {
		t.Fatal("Session.Put test failed")
	}
}

func TestSessionPatch(t *testing.T) {
	ts := newTestSessionServer()
	defer ts.Close()

	session := NewSession()
	resp, err := session.Patch(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	text := resp.Text()
	if text != "PATCH" {
		t.Fatal("Session.Patch test failed")
	}
}

func TestSessionDelete(t *testing.T) {
	ts := newTestSessionServer()
	defer ts.Close()

	session := NewSession()
	resp, err := session.Delete(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	text := resp.Text()
	if text != "DELETE" {
		t.Fatal("Session.Delete test failed")
	}
}

func TestSessionHead(t *testing.T) {
	ts := newTestSessionServer()
	defer ts.Close()

	session := NewSession()
	resp, err := session.Head(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	cookies := resp.Cookies
	if cookies[0].Name != "HEAD" {
		t.Fatal("Session.Head test failed")
	}
}

func TestSessionCookieJar(t *testing.T) {
	ts := newTestSessionServer()
	defer ts.Close()

	session := NewSession()
	_, err := session.Get(ts.URL + "/setCookie")
	if err != nil {
		t.Fatal(err)
	}
	resp, err := session.Get(ts.URL + "/getCookie")
	if err != nil {
		t.Fatal(err)
	}
	if resp.Text() != "key=value" {
		t.Fatal("Session.CookieJar failed.")
		return
	}
}

func TestSessionSetCookie(t *testing.T) {
	ts := newTestSessionServer()
	defer ts.Close()

	session := NewSession()
	cookie := NewCookies("key", "value")
	session.SetCookies(ts.URL, cookie)
	resp, err := session.Get(ts.URL + "/getCookie")
	if err != nil {
		t.Fatal(err)
	}
	if resp.Text() != "key=value" {
		t.Fatal("Session.SetCookies() failed.")
		return
	}
}

func TestSessionCookies(t *testing.T) {
	ts := newTestSessionServer()
	defer ts.Close()

	session := NewSession()
	_, err := session.Get(ts.URL + "/setCookie")
	if err != nil {
		t.Fatal(err)
	}
	cookies := session.Cookies(ts.URL)
	if cookies[0].Name != "key" {
		t.Fatal("Session.Cookies() failed.")
	}
}
