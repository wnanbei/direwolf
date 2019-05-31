package direwolf

import (
	"fmt"
	"testing"
)

func TestStringSliceMap(t *testing.T) {
	c := NewCookies(
		"key1", "key2",
		"key3", "key4",
	)
	// s.Add("hello", "world")
	// s.Add("hello", "world")
	// s.Add("hello", "world")
	// s.Add("hello", "world")

	fmt.Println(c.URLEncode())
}
