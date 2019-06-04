package direwolf

import (
	"fmt"
	"testing"
)

func first1() error {
	err := first2()
	return MakeErr(err, "next", "first1")
}

func first2() error {
	return MakeErr(nil, "Error", "wrong")
}

func TestError(t *testing.T) {
	err := first1()
	fmt.Println(err.Error())
	// fmt.Println(err.Msg)
}
