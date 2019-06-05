package direwolf

import (
	"fmt"
	"runtime"
)

// var HTTPError = "HTTPError"
// var EncodeError = "EncodeError"
// var TypeError = "TypeError"

// ErrorStack is a wrapped error type, contain stack info, like file name and code line
type ErrorStack struct {
	Prev error
	Msg  string
	File string
	Line int
}

// Error method made ErrorStack type contain the error stack.
func (e ErrorStack) Error() string {
	if e.Prev == nil {
		return fmt.Sprintf("%s:%d\n%s", e.File, e.Line, e.Msg)
	}
	return fmt.Sprintf("%s:%d\n%s\n%v", e.File, e.Line, e.Msg, e.Prev)
}

// MakeErrorStack used to create ErrorStack type, prev is the previous error.
func MakeErrorStack(prev error, msg string) error {
	_, file, line, _ := runtime.Caller(1)
	return &ErrorStack{
		Prev: prev,
		Msg:  msg,
		File: file,
		Line: line,
	}
}

// Error combined by Name field and ErrorStack. So it have name and stack info.
type Error struct {
	Name string
	ErrorStack
}

// Error method made Error type contain the error stack.
func (e Error) Error() string {
	if e.Prev == nil {
		return fmt.Sprintf("%s:%d - %s\n%s", e.File, e.Line, e.Name, e.Msg)
	}
	return fmt.Sprintf("%s:%d - %s\n%s\n%v", e.File, e.Line, e.Name, e.Msg, e.Prev)
}

// MakeError used to create Error type, prev is the previous error.
func MakeError(prev error, name, msg string) error {
	_, file, line, _ := runtime.Caller(1)
	return &Error{
		Name: name,
		ErrorStack: ErrorStack{
			Prev: prev,
			Msg:  msg,
			File: file,
			Line: line,
		},
	}
}
