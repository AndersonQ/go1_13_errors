package myerrors

import (
	"errors"
	"fmt"
)

func ExampleMyError_Wrap() {
	e := Wrap(errors.New("a error"), "wrapping msg")

	fmt.Println(e)
	// Output: wrapping msg: a error
}

func ExampleMyError_MultipleWrap() {
	e := Wrap(Wrap(errors.New("a error"), "wrapping msg"), "another wrapping msg")

	fmt.Println(e)
	// Output: another wrapping msg: wrapping msg: a error
}
