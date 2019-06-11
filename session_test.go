package direwolf

import (
	"fmt"
	"testing"
)

type myType struct {
	m map[string]string
}

func TestSession(t *testing.T) {
	session := NewSession()
	session.Cookies.New(
		"hello", "world",
		"key", "value",
	)
	resp, err := session.Get("https://httpbin.org/get")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Text())

	cookies := NewCookies(
		"xxx", "yyy",
	)
	resp, err = session.Get("https://httpbin.org/get", cookies)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Text())

	resp, err = session.Get("https://httpbin.org/get")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Text())
}
