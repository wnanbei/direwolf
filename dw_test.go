package direwolf

import (
	"testing"
)

func TestHTTP(t *testing.T) {
	h := Headers{
		"aaa": {"bbb", "ccc", "bbb"},
	}
	req := new(Request)
	req.setHeader(h)
	t.Log(req.Headers)
}

func TestGet(t *testing.T) {
	h := Headers{
		"aaa": {"bbb", "ccc", "bbb"},
	}
	Get("http://httpbin.org/get", h)
}
