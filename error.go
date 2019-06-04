package direwolf

import (
	"fmt"
	"runtime"
)

// Error is the wrapped error type, contain file name, file line
type Error struct {
	Prev error
	Name string
	Msg  string
	File string
	Line int
}

// Error method made Error type contain the error stack.
func (e Error) Error() string {
	if e.Prev == nil {
		return fmt.Sprintf("%s:%d-%s\n%s", e.File, e.Line, e.Name, e.Msg)
	}
	return fmt.Sprintf("%s:%d-%s\n%s\n%v", e.File, e.Line, e.Name, e.Msg, e.Prev)
}

// MakeErr used to create Error type, prev is the previous error.
func MakeErr(prev error, name, msg string) error {
	_, file, line, _ := runtime.Caller(1)
	return Error{
		Prev: prev,
		Name: name,
		Msg:  msg,
		File: file,
		Line: line,
	}
}

// func MakeHTTPErr(prev error, msg string) error {
// 	_, file, line, _ := runtime.Caller(1)
// 	return Error{
// 		Prev: prev,
// 		Name: "HTTPError",
// 		Msg:  msg,
// 		File: file,
// 		Line: line,
// 	}
// }
