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
