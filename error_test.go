package direwolf

import (
	"fmt"
	"testing"
)

func f1() error {
	err := f2()
	if err != nil {
		return MakeErrorStack(err, "msg")
	}
	return nil
}

func f2() error {
	return MakeError(nil, "HTTPError", "msg")
}

func TestError(t *testing.T) {
	err := f1()
	fmt.Println(err)
}
