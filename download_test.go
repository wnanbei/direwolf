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
