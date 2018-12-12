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

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

type sages struct {
	Name  string
	Motto string
}

type musicians struct {
	ShowName string
	Band     string
}

type people struct {
	SageSliceStruct     []sages
	MusicianSliceStruct []musicians
}

// //////////////////////////////////////////////////////////////////////////////////
// MAIN FUNCTION
// //////////////////////////////////////////////////////////////////////////////////

func main() {

	// //////////////////////////////////////////////////////////////////////////////////
	// INFO FOR SAGES STRUCT
	// //////////////////////////////////////////////////////////////////////////////////

	buddha := sages{
		Name:  "Buddha",
		Motto: "The belief of no beliefs.",
	}

	mlk := sages{
		Name:  "Martin Luther King",
		Motto: "Hated never ceases with hatred but with love alone is healed.",
	}

	gandhi := sages{
		Name:  "Gandhi",
		Motto: "Be the change.",
	}

	jesus := sages{
		Name:  "Jesus",
		Motto: "Love all.",
	}

	muhammad := sages{
		Name:  "Muhammad",
		Motto: "To overcome evil with good is good, to resist evil by evil is evil.",
	}

	// //////////////////////////////////////////////////////////////////////////////////
	// INFO FOR MUSICIANS STRUCT
	// //////////////////////////////////////////////////////////////////////////////////

	corpseGrinder := musicians{
		ShowName: "Corpse Grinder",
		Band:     "Cannabil Corpse",
	}

	erikRutan := musicians{
		ShowName: "Erik Rutan",
		Band:     "Hate Eternal",
	}

	muhammadSuicmez := musicians{
		ShowName: "Muhammad Suicmez",
		Band:     "Necrophagist",
	}

	floMounier := musicians{
		ShowName: "Flo Mounier",
		Band:     "Cryptopsy",
	}

	// //////////////////////////////////////////////////////////////////////////////////
	// INFO FOR PEOPLE STRUCT
	// //////////////////////////////////////////////////////////////////////////////////

	people := people{
		SageSliceStruct:     []sages{buddha, mlk, gandhi, jesus, muhammad},
		MusicianSliceStruct: []musicians{corpseGrinder, erikRutan, muhammadSuicmez, floMounier},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", people)
	if err != nil {
		log.Fatalln(err)
	}

}
