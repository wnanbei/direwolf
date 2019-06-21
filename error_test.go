package direwolf

import (
	"testing"
)

func TestError(t *testing.T) {
	err := MakeError(nil, HTTPError, "testError")
	t.Log(MakeErrorStack(err, "testing"))
}
