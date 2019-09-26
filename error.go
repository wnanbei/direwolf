package direwolf

import (
	"fmt"
	"runtime"
	"time"
)

const (
	// HTTPError means some thing wrong while request.
	HTTPError = "HTTPError"
	// NewRequestError means create request method field.
	NewRequestError = "NewRequestError"
	// RequestBodyError means request body can`t exists with post form.
	RequestBodyError = "RequestBodyError"
	// ProxyURLError means that proxy url was wrong.
	ProxyURLError = "ProxyURLError"
	// RedirectError means over the max number of redirect.
	RedirectError = "RedirectError"
	// URLError means url is wrong.
	URLError = "URLError"
)

type Error struct {
	// wrapped error
	err      error
	msg      string
	// file path and name
	file     string
	fileLine int
	time     string
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s - %s:%d\n%s\n\n%s", e.time, e.file, e.fileLine, e.msg, e.err.Error())
}

func (e *Error) Unwrap() error {
	if e.err != nil {
		return e.err
	}
	return nil
}

// stack will return the error stack information with a string.
//func (e *Error) stack() string {
//	return fmt.Sprintf("[%s] - %s:%d - %s\n%s", e.time, e.file, e.fileLine, e.msg, ErrStack(e.err))
//}

// WrapError will wrap a error with some information: filename, line, time.
func WrapError(err error, msg string) error {
	_, file, line, _ := runtime.Caller(1)
	return &Error{
		err:      err,
		msg:      msg,
		file:     file,
		fileLine: line,
		time:     time.Now().Format("2006-01-02 15:04:05"),
	}
}

// ErrStack will return error stack information.
//func ErrStack(err error) string {
//	var prevErr *Error
//	if errors.As(err, &prevErr) {
//		return prevErr.stack()
//	}
//	return err.Error()
//}



// ErrorStack is a wrapped error type, contain stack info, like file name and code line.

// Error method made ErrorStack type contain the error stack.
// MakeErrorStack used to create ErrorStack type, prev is the previous error.

// Error combined by Name field and ErrorStack. So it have name and stack info.

// Error method made Error type contain the error stack.

// MakeError used to create Error type, prev is the previous error.
