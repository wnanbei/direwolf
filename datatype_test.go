package direwolf

import (
	"fmt"
	"testing"
)

func TestStringSliceMap(t *testing.T) {
	r := NewRequestSetting("Get", "https://www.baidu.com")
	c := NewCookies(
		"key1", "key2",
		"key3", "key4",
	)
	r.Cookies = c
	if r.Cookies.data != nil {
		fmt.Println(c.URLEncode())
	}
}

func TestBody(t *testing.T) {
	body := Body("ddd")
	req := NewRequestSetting("get", "http")
	req.Body = body
	t.Log(req.Body)
}
