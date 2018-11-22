package main

/*
- Initialize a SLICE of int using a composite literal
- print out the slice
- range over the slice printing out just the index
- range over the slice printing out both the index and the value
*/

import (
	"fmt"
)

func main() {
	favNum := []int{42, 3, 0, 1, 6, 9, 13}

	fmt.Println(favNum)

	for i := range favNum {
		fmt.Println(i)
	}

	for i, num := range favNum {
		fmt.Println(i, num)
	}
}
