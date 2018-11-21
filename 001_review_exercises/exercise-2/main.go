package main

/*
- create a struct that holds person fields
- create a struct that holds secret agent fields and embeds person type
- attach a method to person: pSpeak
- attach a method to secret agent: saSpeak
- create a variable of type person
- create a variable of type secret agent
- print a field from person
- run pSpeak attached to the variable of type person
- print a field from secret agent
- run saSpeak attached to the variable of type secret agent
- run pSpeak attached to the variable of type secret agent
- create an interface type that both person and secretAgent implement
- declare a func with a parameter of the interface’s type
- call that func in main and pass in a value of type person
- call that func in main and pass in a value of type secretAgent
*/

import "fmt"

// //////////////////////////////////////////////////////////////////////////////////
// DATA DECLARATIONS
// //////////////////////////////////////////////////////////////////////////////////

type person struct {
	fName string
	lName string
}

type secretAgent struct {
	person
	agency string
}

type homosapien interface {
	speak()
}

// //////////////////////////////////////////////////////////////////////////////////
// PACKAGE FUNCTIONS
// //////////////////////////////////////////////////////////////////////////////////

func (p person) speak() {
	fmt.Println(p.fName)
}

func (sa secretAgent) speak() {
	fmt.Println(sa.agency, sa.fName, sa.lName)
}

func consciousHuman(h homosapien) {
	fmt.Println(h)
}

// //////////////////////////////////////////////////////////////////////////////////
// MAIN FUNCTION
// //////////////////////////////////////////////////////////////////////////////////

func main() {
	dndChar := person{fName: "Nas", lName: "Lókë"}

	saDndChar := secretAgent{
		person: person{fName: "Rogar", lName: "The Rogue"},
		agency: "Kingsguard:",
	}

	dndChar.speak()
	saDndChar.speak()
	saDndChar.speak()

	consciousHuman(saDndChar)
	consciousHuman(dndChar)
}
