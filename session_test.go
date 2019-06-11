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
	resp, err := session.Post("https://httpbin.org/post")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Text())
}
