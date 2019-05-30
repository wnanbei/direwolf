package direwolf

import (
	"fmt"
	"testing"
)

func TestStringSliceMap(t *testing.T) {
	s := StringSliceMap{}.New(
		"key1", "key2",
		"key3", "key4", "key5",
	)
	// s.Add("hello", "world")
	// s.Add("hello", "world")
	// s.Add("hello", "world")
	// s.Add("hello", "world")

	fmt.Println(s.URLEncode())
}
