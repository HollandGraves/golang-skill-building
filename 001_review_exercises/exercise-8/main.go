package main

/*
- Using the short declaration operator, create a variable with the identifier “s”
- assign “i'm sorry dave i can't do that” to “s”
- Print “s”.
- Print “s” converted to a slice of byte.
- Print “s” converted to a slice of byte and then converted back to a string.
- Using slicing, print just “i’m sorry dave”
- Using slicing, print just “dave i can’t”
- Using slicing, print just “can’t do that”
- print every letter of “s” with one rune (character) on each line
*/

import (
	"fmt"
)

// //////////////////////////////////////////////////////////////////////////////////
// DATA DECLARATIONS
// //////////////////////////////////////////////////////////////////////////////////

// //////////////////////////////////////////////////////////////////////////////////
// PACKAGE FUNCTIONS
// //////////////////////////////////////////////////////////////////////////////////

// //////////////////////////////////////////////////////////////////////////////////
// MAIN FUNCTION
// //////////////////////////////////////////////////////////////////////////////////

func main() {
	s := "i'm sorry dave i can't do that"
	fmt.Println(s)
	fmt.Println([]byte(s))
	fmt.Println(string([]byte(s)))
	// turning string into byte slice, use specific range of byte slice,
	// then convert specific range back to string
	fmt.Println(string([]byte(s)[:14]))
	fmt.Println(string([]byte(s)[10:22]))
	fmt.Println(string([]byte(s)[17:30]))

	for i := range s {
		fmt.Println(string([]byte(s)[i : i+1]))
	}

	// alternate way
	for _, v := range s {
		fmt.Println(string(v))
	}
}
