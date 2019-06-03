package direwolf

import (
	"fmt"
	"runtime"
)

type Error struct {
	Prev error
	Err  string
	Msg  string
	File string
	Line int
}

func (e Error) Error() string {
	if e.Prev == nil {
		return fmt.Sprintf("%s @ %d: %s: %s", e.File, e.Line, e.Err, e.Msg)
	}
	return fmt.Sprintf("%s @ %d: %s: %s\n%v", e.File, e.Line, e.Err, e.Msg, e.Prev)
}

func MakeErr(prev error, code string, format string, args ...interface{}) error {
	_, file, line, _ := runtime.Caller(1)
	return Error{
		Prev: prev,
		Err:  code,
		Msg:  fmt.Sprintf(format, args...),
		File: file,
		Line: line,
	}
}
