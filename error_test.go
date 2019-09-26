package direwolf

import (
	"errors"
	"fmt"
	"testing"
)

func TestError(t *testing.T) {
	err1 := errors.New("first error")
	err2 := WrapError(err1, "second testing")
	err3 := WrapError(err2, "third testing")
	fmt.Println(err3)
}
