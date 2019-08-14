package direwolf

import (
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func startWebserver() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Millisecond * 50)
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	go http.ListenAndServe(":8777", nil)
}

func startLoadTest() {
	count := 0
	for {
		resp, err := Get("http://localhost:8777/")
		if err != nil {
			panic(fmt.Sprintf("Got error: %v", err))
		}
		resp.Text()
		log.Printf("Finished GET request #%v", count)
		count++
	}
}

func TestSession(t *testing.T) {
	// start a webserver in a goroutine
	// s = NewSession()
	// s.transport.MaxIdleConnsPerHost = 100
	// s.transport.DisableKeepAlives = true

	startWebserver()
	for i := 0; i < 100; i++ {
		go startLoadTest()
	}
	time.Sleep(time.Second * 2400)
}

func TestSessionGet(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// check method is GET before going to check other features
		if r.Method != "GET" {
			t.Fatalf("Expected method %q; got %q", "GET", r.Method)
		}
		if r.URL.Path == "/test" {
			w.Write([]byte("passed"))
		}
	}))
	defer ts.Close()

	session := NewSession()
	resp, err := session.Get(ts.URL + "/test")
	if err != nil {
		t.Fatal(err)
	}
	text := resp.Text()
	if text != "passed" {
		t.Fatal("response was wrong, not", text)
	}
	t.Log("Session.Get test passed")
}

func TestSessionPost(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// check method is GET before going to check other features
		if r.Method != "POST" {
			t.Fatalf("Expected method %q; got %q", "Post", r.Method)
		}
		body, _ := ioutil.ReadAll(r.Body)
		if string(body) != "key=value" {
			w.Write([]byte("Post body failed."))
		} else {
			w.Write([]byte("passed"))
		}
	}))
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
	if text != "passed" {
		t.Fatal("response was wrong,", text)
	}

	body := Body("key=value")
	resp2, err := session.Post(ts.URL+"/test", body)
	if err != nil {
		t.Fatal(err)
	}
	text2 := resp2.Text()
	if text2 != "passed" {
		t.Fatal("response was wrong,", text)
	}

	t.Log("Session.Post test passed")
}
