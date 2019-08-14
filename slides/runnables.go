package main

import (
	"errors"
	"fmt"
)

func main() {
	// start_fmt OMIT
	err0 := errors.New("my error")                                    // HL
	err1 := fmt.Errorf("1s wrapping my error with Errorf: %w", err0)  // HL
	err2 := fmt.Errorf("2nd wrapping my error with Errorf: %w", err1) // HL

	fmt.Println(err0)
	fmt.Println(err1)
	fmt.Println(err2)
	// end_fmt OMIT
}
