package direwolf

import (
	"testing"
)

func TestGet(t *testing.T) {
	ts := newTestSessionServer()
	defer ts.Close()

	resp, err := Get(ts.URL + "/test")
	if err != nil {
		t.Fatal(err)
	}
	text := resp.Text()
	if text != "GET" {
		t.Fatal("Get test failed")
	}
}

func TestPost(t *testing.T) {
	ts := newTestSessionServer()
	defer ts.Close()

	postForm := NewPostForm(
		"key", "value",
	)
	resp, err := Post(ts.URL+"/test", postForm)
	if err != nil {
		t.Fatal(err)
	}
	text := resp.Text()
	if text != "POST" {
		t.Fatal("Post test failed")
	}

	body := Body("key=value")
	resp2, err := Post(ts.URL+"/test", body)
	if err != nil {
		t.Fatal(err)
	}
	text2 := resp2.Text()
	if text2 != "POST" {
		t.Fatal("Post test failed")
	}
}

func TestPut(t *testing.T) {
	ts := newTestSessionServer()
	defer ts.Close()

	resp, err := Put(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	text := resp.Text()
	if text != "PUT" {
		t.Fatal("Put test failed")
	}
}

func TestPatch(t *testing.T) {
	ts := newTestSessionServer()
	defer ts.Close()

	resp, err := Patch(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	text := resp.Text()
	if text != "PATCH" {
		t.Fatal("Patch test failed")
	}
}

func TestDelete(t *testing.T) {
	ts := newTestSessionServer()
	defer ts.Close()

	resp, err := Delete(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	text := resp.Text()
	if text != "DELETE" {
		t.Fatal("Delete test failed")
	}
}

func TestHead(t *testing.T) {
	ts := newTestSessionServer()
	defer ts.Close()

	resp, err := Head(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	cookies := resp.Cookies
	if cookies[0].Name != "HEAD" {
		t.Fatal("Head test failed")
	}
}

func TestRequest(t *testing.T) {
	ts := newTestSessionServer()
	defer ts.Close()

	req := NewRequest("Get", ts.URL+"/test")
	resp, err := Send(req)
	if err != nil {
		t.Fatal(err)
	}
	text := resp.Text()
	if text != "GET" {
		t.Fatal("Request test failed")
	}
}

func TestSendCookie(t *testing.T) {
	ts := newTestSessionServer()
	defer ts.Close()

	cookies := NewCookies(
		"name", "direwolf",
	)
	resp, err := Get(ts.URL+"/getCookie", cookies)
	if err != nil {
		t.Fatal(err)
	}
	if resp.Text() != "name=direwolf" {
		t.Fatal("request cookies test failed")
	}
}

func TestSendHeaders(t *testing.T) {
	ts := newTestSessionServer()
	defer ts.Close()

	headers := NewHeaders(
		"User-Agent", "direwolf",
	)
	resp, err := Get(ts.URL+"/getHeader", headers)
	if err != nil {
		t.Fatal(err)
	}
	text := resp.Text()
	if text != "direwolf" {
		t.Fatal("request headers test failed")
	}
}

func TestSendParams(t *testing.T) {
	ts := newTestSessionServer()
	defer ts.Close()

	params := NewParams(
		"key", "value",
	)
	resp, err := Get(ts.URL+"/getParams", params)
	if err != nil {
		t.Fatal(err)
	}
	if resp.Text() != "value" {
		t.Fatal("request params test failed")
	}
}
