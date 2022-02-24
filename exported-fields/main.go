package main

import (
	"fmt"

	"github.com/codegold79/go-playground/exported-fields/nulls"
)

func main() {
	var s *nulls.NullString

	// s is still nil, so fields cannot be updated.
	s.Set("abc", "hidden secret")
	fmt.Printf("s type: %T, val: %#[1]v\n\n", s)

	s2 := &nulls.NullString{}
	s2.Set("abc", "hidden secret")
	// The hidden value can be printed, but without a getter on NullString, it
	// cannot programmatically be accessed.
	fmt.Printf("s type: %T, val: %#[1]v\n", s2)
}
