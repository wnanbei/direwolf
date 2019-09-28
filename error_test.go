package direwolf

import (
	"errors"
	"testing"
)

func TestError(t *testing.T) {
	err1 := ErrRequestBody
	err2 := WrapErr(err1, "second testing")
	err3 := WrapErrf(err2, "==%s==", "third testing")

	if !errors.Is(err3, ErrRequestBody) {
		t.Fatal("Test errors.Is failed.")
	}
}

func TestRedirectError(t *testing.T) {
	red := RedirectNum(3)
	_, err := Get("http://httpbin.org/redirect/4", red)
	if err != nil {
		var eType *RedirectError
		if !errors.As(err, &eType) {
			t.Fatal("Test TestRedirectError failed.")
		}
	} else {
		t.Fatal("Test TestRedirectError failed.")
	}
}
