package direwolf

import (
	"errors"
	"fmt"
	"runtime"
	"strconv"
	"time"
)

var (
	ErrRequestBody = errors.New("request body can`t coexists with PostForm")
)

type RedirectError struct {
	RedirectNum int
}

func (e *RedirectError) Error() string {
	return "exceeded the maximum number of redirects: " + strconv.Itoa(e.RedirectNum)
}

type Error struct {
	// wrapped error
	err error
	msg string
	// file path and name
	file     string
	fileLine int
	time     string
}

func (e *Error) Error() string {
	_, ok := e.err.(interface {
		Unwrap() error
	})
	if ok {
		return fmt.Sprintf("%s - %s:%d\n%s\n%s", e.time, e.file, e.fileLine, e.msg, e.err.Error())
	}
	return fmt.Sprintf("%s - %s:%d\n%s\n\n%s\n", e.time, e.file, e.fileLine, e.msg, e.err.Error())
}

func (e *Error) Unwrap() error {
	if e.err != nil {
		return e.err
	}
	return nil
}

// WrapErr will wrap a error with some information: filename, line, time and some message.
func WrapErr(err error, msg string) error {
	_, file, line, _ := runtime.Caller(1)
	return &Error{
		err:      err,
		msg:      msg,
		file:     file,
		fileLine: line,
		time:     time.Now().Format("2006-01-02 15:04:05"),
	}
}

// WrapErr will wrap a error with some information: filename, line, time and some message.
// You can format message of error.
func WrapErrf(err error, format string, args ...interface{}) error {
	msg := fmt.Sprintf(format, args...)
	_, file, line, _ := runtime.Caller(1)
	return &Error{
		err:      err,
		msg:      msg,
		file:     file,
		fileLine: line,
		time:     time.Now().Format("2006-01-02 15:04:05"),
	}
}
