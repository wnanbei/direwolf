package direwolf

import (
	"testing"
	"time"
)

func TestMergeHeaders(t *testing.T) {
	h1 := NewHeaders(
		"hello", "world",
	)
	h2 := NewHeaders(
		"key", "value",
		"hello", "shit",
	)
	h := mergeHeaders(h1, h2)
	t.Log(h)
}

func TestTimeout(t *testing.T) {
	a := 15

	mytime := time.Duration(a) * time.Second
	t.Log(mytime)
}

func TestRedirect(t *testing.T) {
	headers := NewHeaders(
		"User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.108 Safari/537.36",
	)
	params := NewParams(
		"keyxxx", "南北",
	)
	cookies := NewCookies(
		"hello", "world",
	)
	resp, err := Get(
		"https://httpbin.org/2",
		headers,
		params,
		cookies,
		RedirectNum(2),
	)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(resp.Text())
	}
}
