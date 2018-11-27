package main

import (
	"log"
	"os"
	"text/template"
)

// //////////////////////////////////////////////////////////////////////////////////
// DATA DECLARATIONS
// //////////////////////////////////////////////////////////////////////////////////

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

// //////////////////////////////////////////////////////////////////////////////////
// MAIN FUNCTION
// //////////////////////////////////////////////////////////////////////////////////

func main() {
	buddha := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs.",
	}

	mlk := sage{
		Name:  "Martin Luther King",
		Motto: "Hated never ceases with hatred but with love alone is healed.",
	}

	gandhi := sage{
		Name:  "Gandhi",
		Motto: "Be the change.",
	}

	jesus := sage{
		Name:  "Jesus",
		Motto: "Love all.",
	}

	muhammad := sage{
		Name:  "Muhammad",
		Motto: "To overcome evil with good is good, to resist evil by evil is evil.",
	}

	sages := []sage{buddha, mlk, gandhi, jesus, muhammad}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", sages)
	if err != nil {
		log.Fatalln("Error:", err)
	}
}
