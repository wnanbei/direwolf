package direwolf

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

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
