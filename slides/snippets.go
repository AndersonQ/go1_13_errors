// package main

// start_unwrap OMIT
// Unwrap returns the result of calling the Unwrap method on err, if err's
// type contains an Unwrap method returning error.
// Otherwise, Unwrap returns nil.
type Wrapper interface {
	Unwrap() error // HL
}
// end_unwrap OMIT

// start_is OMIT
package errors

// Is reports whether any error in err's chain matches target.
//
// The chain consists of err itself followed by the sequence of errors obtained by
// repeatedly calling Unwrap.
//
// An error is considered to match a target if it is equal to that target or if
// it implements a method Is(error) bool such that Is(target) returns true.
func Is(err, target error) bool // HL
// end_is OMIT


// start_as OMIT
package errors

// As finds the first error in err's chain that matches target, and if so, sets
// target to that error value and returns true.
//
// An error matches target if the error's concrete value is assignable to the value
// pointed to by target, or if the error has a method As(interface{}) bool such that
// As(target) returns true. In the latter case, the As method is responsible for
// setting target.
//
// As will panic if target is not a non-nil pointer to either a type that implements
// error, or to any interface type. As returns false if err is nil.
func As(err error, target interface{}) bool // HL

// end_as OMIT

// start_go1.13 OMIT

u, ok := err.(interface { Unwrap() error})

x, ok := err.(interface { Is(error) bool })

x, ok := err.(interface { As(interface{}) bool })

// end_go1.13 OMIT

// start_unwrap_go1.13 OMIT
// Unwrap returns the result of calling the Unwrap method on err, if err's
// type contains an Unwrap method returning error.
// Otherwise, Unwrap returns nil.
	Unwrap() error // HL
// end_unwrap_go1.13 OMIT
