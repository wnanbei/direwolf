package direwolf

import (
	"testing"
)

//func TestMini(t *testing.T) {
//	session := NewSession()
//	_, err := session.Get("http://httpbin.org/cookies/set/key/value")
//	if err != nil {
//		return
//	}
//	fmt.Println("first step")
//	cookies := session.Cookies("http://httpbin.org")
//	fmt.Println(cookies)
//}

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
	t.Log("Get test passed")
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
	t.Log("Post test passed")
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
	t.Log("Put test passed")
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
	t.Log("Patch test passed")
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
	t.Log("Delete test passed")
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
	t.Log("Head test passed")
}

func TestRequest(t *testing.T) {
	ts := newTestSessionServer()
	defer ts.Close()

	req := NewRequestSetting("Get", ts.URL+"/test")
	resp, err := Request(req)
	if err != nil {
		t.Fatal(err)
	}
	text := resp.Text()
	if text != "GET" {
		t.Fatal("Request test failed")
	}
	t.Log("Request test passed")
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
	text := resp.Text()
	if text != "name=direwolf" {
		t.Fatal("request cookies test failed")
	}
	t.Log("request cookies test passed")
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
	t.Log("request headers test passed")
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
	text := resp.Text()
	if text != "value" {
		t.Fatal("request headers test failed")
	}
	t.Log("request headers test passed")
}
