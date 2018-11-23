package main

/*
- Create a new type called person which is STRUCT that stores fName and lName
- Using the STRUCT “person” and a composite literal create a value of type person and assign it to a variable with the identifier “p1”
- print out “p1”
- print out just the field fName for “p1”
- Take the STRUCT “person” and add a field called “favFood” that stores a slice of string
- for the variable “p1” use a composite literal to add values to the field favFood
- print out the values in favFood
- also print out the values for “favFood” by using a for range loop
- Add a method to type “person” with the identifier “walk”
- Have the func return this string: <person’s first name> is walking
*/

import "fmt"

// //////////////////////////////////////////////////////////////////////////////////
// DATA DECLARATIONS
// //////////////////////////////////////////////////////////////////////////////////

type person struct {
	fName   string
	lName   string
	favFood []string
}

// //////////////////////////////////////////////////////////////////////////////////
// PACKAGE FUNCTIONS
// //////////////////////////////////////////////////////////////////////////////////

func (p *person) walk() {
	fmt.Printf("%s is walking\n", p.fName)

	return
}

// //////////////////////////////////////////////////////////////////////////////////
// MAIN FUNCTION
// //////////////////////////////////////////////////////////////////////////////////

func main() {
	p1 := person{
		fName: "Nas",
		lName: "Lóckë",
		favFood: []string{
			"Ramen",
			"Pork",
			"Kombucha",
		},
	}

	fmt.Println(p1)
	fmt.Println(p1.fName)
	fmt.Println(p1.favFood)

	for _, food := range p1.favFood {
		fmt.Println(food)
	}

	p1.walk()

}
