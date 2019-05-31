package direwolf

import (
	"fmt"
	"testing"
)

func TestStringSliceMap(t *testing.T) {
	r := NewRequestSetting()
	c := NewCookies(
		"key1", "key2",
		"key3", "key4",
	)
	r.Cookies = c
	// s.Add("hello", "world")
	// s.Add("hello", "world")
	// s.Add("hello", "world")
	// s.Add("hello", "world")
	if r.Cookies.data != nil {
		fmt.Println(c.URLEncode())
	}
}
