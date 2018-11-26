package main

/*
- Create a new type called “gator”, the underlying type of “gator” is an int
- Using var, declare an identifier “g1” as type gator
- Assign a value to “g1”
- Print out “g1”
- Print the type of “g1” using fmt.Printf(“%T\n”, g1)
- Using var, declare an identifier “x” as type int
- Print out “x”
- Print the type of “x” using fmt.Printf(“%T\n”, x)
- Can you assign “g1” to “x”? Why or why not?
- Now add a method to type gator with this signature greeting()
- have it print “Hello, I am a gator”
- Call the greeting() method.
- create another type called “flamingo” with the underlying type of “flamingo” bool
- Give “flamingo” a method with this signature greeting()
- and have it print “Hello, I am pink and beautiful and wonderful.”
- create a new type “swampCreature”
- The underlying type of “swapCreature” is interface
- The behavior which the “swampCreature” interface defines is such that any type which has this method greeting()
- Create a func called “bayou” which takes a value of type “swampCreature” as an argument
- Have this func print out the greeting for whatever “swampCreature” might be passed in
*/

import (
	"fmt"
)

// //////////////////////////////////////////////////////////////////////////////////
// DATA DECLARATIONS
// //////////////////////////////////////////////////////////////////////////////////

type gator int

type flamingo bool

type swampCreature interface {
	greeting()
}

// //////////////////////////////////////////////////////////////////////////////////
// PACKAGE FUNCTIONS
// //////////////////////////////////////////////////////////////////////////////////

func (g gator) greeting() {
	fmt.Println("Hello, I am a gator")
}

func (f flamingo) greeting() {
	fmt.Println("Hello, I am pink and beautiful and wonderful")
}

func bayou(sw swampCreature) {
	sw.greeting()
}

// //////////////////////////////////////////////////////////////////////////////////
// MAIN FUNCTION
// //////////////////////////////////////////////////////////////////////////////////

func main() {
	var g1 gator

	g1 = 7
	fmt.Println(g1)
	fmt.Printf("%T\n", g1)

	var x int
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	// You cannont assign x = g1 because x has a zero value of 0 and type int
	// while g1 has a value of 7 and type gator

	g1.greeting()

	var f1 flamingo
	f1 = true

	bayou(g1)
	bayou(f1)
}
