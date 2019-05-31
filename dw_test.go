package direwolf

import (
	"fmt"
	"testing"
)

func TestHTTP(t *testing.T) {
	headers := Headers{
		"User-Agent": {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.108 Safari/537.36"},
	}
	params := NewParams(
		"keyxxx", "valuexxx",
	)
	cookies := NewCookies(
		"hello", "world",
	)
	var proxy Proxy = "http://127.0.0.1:12333"
	resp := Get("http://httpbin.org/get", headers, params, proxy, cookies)
	fmt.Println(resp.Text())
}

type myType map[string][]string

func (m myType) new(slice ...[]string) {
	m = make(map[string][]string, 10)
	for _, kv := range slice {
		m[kv[0]] = append(m[kv[0]], kv[1])
	}
}

// func (m myType) all() map[string][]string {
// 	return m
// }

type cookie struct {
	myType
}

func newCookie(slice ...[]string) cookie {
	var c = cookie{}
	c.new(slice...)
	return c
}

// func x(c cookie) {
// 	for key, value := range c {
// 		fmt.Println(key)
// 		fmt.Println(value)
// 	}
// }

func TestCookie(t *testing.T) {
	c := newCookie(
		[]string{"key", "value"},
		[]string{"key2", "value"},
		[]string{"key3", "value"},
	)
	// b := c.new(
	// 	[]string{"key", "value"},
	// 	[]string{"key2", "value"},
	// 	[]string{"key3", "value"},
	// )
	fmt.Println(c)
}
