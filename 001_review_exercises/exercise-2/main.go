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

// //////////////////////////////////////////////////////////////////////////////////
// PACKAGE FUNCTIONS
// //////////////////////////////////////////////////////////////////////////////////

func (p *person) pSpeak() {
	fmt.Println(p.fName)
}

func (sa *secretAgent) saSpeak() {
	fmt.Println(sa.agency, sa.fName, sa.lName)
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

	dndChar.pSpeak()
	saDndChar.saSpeak()
	saDndChar.pSpeak()
}
