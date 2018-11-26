package main

import (
	"log"
	"os"
	"text/template"
)

// //////////////////////////////////////////////////////////////////////////////////
// DATA DECLARATIONS
// //////////////////////////////////////////////////////////////////////////////////

// stores pointer (aka container type that will hold all the templates)
var tpl *template.Template

// function that has priority at run time
// is used for initializing variables and many other things
func init() {
	/*
		the reason we use .Must() is so that .Must() can error check
		and return ONLY type *template.Template, that way we don't
		have to extraneously carry around an error type in package
		level declaration

		the reason we are using .ParseFiles() instead of .ParseGlob()
		is because we only have to parse one file, versus a folder of files
	*/
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

// //////////////////////////////////////////////////////////////////////////////////
// MAIN FUNCTION
// //////////////////////////////////////////////////////////////////////////////////

func main() {
	// this time, instead of passing in nil for the third arg,
	// we are passing in a value of type int. This value will
	// be passed into the template where there is a {{.}}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 42)
	if err != nil {
		log.Fatalln(err)
	}
}
