package main

/*
- Initialize a MAP using a composite literal where the key is a string and the value is an int
- print out the map
- range over the map printing out just the key
- range over the map printing out both the key and the value
*/

import (
	"fmt"
)

func main() {
	m := make(map[string]int)

	// one way to initiallize a map with composite literals
	m["favNum1"] = 13
	m["favNum2"] = 7
	m["favNum3"] = 3

	// another way to initiallize a map with composite literals
	n := map[string]int{"favNum4": 1, "favNum5": 9}

	for key, value := range m {
		fmt.Println("Key: ", key, "Value: ", value)
	}

	for key, value := range n {
		fmt.Println("Key: ", key, "Value: ", value)
	}

}
