package myerrors

import (
	"errors"
	"fmt"
)

type myError struct {
	msg string
	err error
}

func New(msg string) error {
	return errors.New(msg)
}

func (m myError) Error() string {
	return m.msg + ": " + m.err.Error()
}

func Wrap(err error, msg string, args ...interface{}) error {
	return myError{
		msg: fmt.Sprintf(msg, args...),
		err: err,
	}
}

func (m myError) Unwrap() error {
	return m.err
}
