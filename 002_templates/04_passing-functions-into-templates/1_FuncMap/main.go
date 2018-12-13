package main

import (
	"html/template"
	"log"
	"os"
	"strings"

	// testing importing functions from a homebrewed dependency
	"github.com/go-udemy-course-exercises/golang-skill-building/002_templates/04_passing-functions-into-templates/1_FuncMap/supportFuncs"
)

// //////////////////////////////////////////////////////////////////////////////////
// DATA DECLARATIONS
// //////////////////////////////////////////////////////////////////////////////////

// DndChar : contains all D&D character info
type DndChar struct {
	Name      string
	Class     string
	Alignment string
	Alive     bool
}

// DungeonMaster : contains all DM info
type DungeonMaster struct {
	Name       string
	YearsDming int
}

// Dnd : to contain all the data and pass into the template
type Dnd struct {
	DndChars       []DndChar
	DungeonMasters []DungeonMaster
}

// functions must be become a FuncMap type so as to be
// able to be passed and, specifically, harnessed by
// a template.
// the FuncMap type is specifically a map[string]interface{}
// the key "tt" is what will be the key to access the function
// WITHIN THE TEMPLATE
var fm = template.FuncMap{
	"tt": strings.Title,
	"fl": firstLetters,

	// function from homebrewed dependency
	"sfl": supportFuncs.FirstLetters,
}

var tpl *template.Template

/*
	Now, first you must make a new pointer to a template with
	the template.New("") method. It takes an argument of a string
	but since the *template doesn't need a name (that will be
	added at the .ParseGlob("") part), we pass in an empty string

	after creating the pointer to a template with template.New("")
	we add in our function definitions and the identifiers that template
	will be able to use

	finally, we .ParseGlob() or ParseFiles() the name of the template(s)
	we wish to pass into the *template struct
*/

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

// //////////////////////////////////////////////////////////////////////////////////
// PACKAGE FUNCTIONS
// //////////////////////////////////////////////////////////////////////////////////

func firstLetters(s string) string {
	// 1. split the string into single words within a slice of string
	// 2. declare var to hold byte type of each first letter of each word
	// 3. loop through slice of string, append each first letter byte
	//	  to the var slice of byte just declared
	// 4. turn slice of bytes back into string
	// 5. capitalize string
	// 6. return string
	fl := strings.Split(s, " ")

	var firstLetterSlice []byte

	for _, firLet := range fl {
		firstLetterSlice = append(firstLetterSlice, firLet[0])
	}

	fls := string(firstLetterSlice)
	ufls := strings.ToUpper(fls)

	return ufls
}

// //////////////////////////////////////////////////////////////////////////////////
// MAIN FUNCTION
// //////////////////////////////////////////////////////////////////////////////////

func main() {
	// Player characters in my current d&d session
	Nas := DndChar{
		Name:      "nas locke",
		Class:     "bard",
		Alignment: "chaotic neutral",
		Alive:     true,
	}

	Nomini := DndChar{
		Name:      "nomini",
		Class:     "druid",
		Alignment: "neutral",
		Alive:     true,
	}

	PaulBunyan := DndChar{
		Name:      "paul bunyan",
		Class:     "barbarian",
		Alignment: "chaotic neutral",
		Alive:     true,
	}

	Rogar := DndChar{
		Name:      "rogar",
		Class:     "rogue",
		Alignment: "chaotic neutral",
		Alive:     true,
	}

	Jorbjorn := DndChar{
		Name:      "jorbjorn",
		Class:     "barbarian",
		Alignment: "chaotic good",
		Alive:     true,
	}

	// DM in my current d&d session
	Logan := DungeonMaster{
		Name:       "logan thibodeaux",
		YearsDming: 10,
	}

	// final struct to contain and pass all the data into the template
	currentDndSession := Dnd{
		DndChars:       []DndChar{Nas, Nomini, PaulBunyan, Rogar, Jorbjorn},
		DungeonMasters: []DungeonMaster{Logan},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", currentDndSession)
	if err != nil {
		log.Fatal(err)
	}

}
