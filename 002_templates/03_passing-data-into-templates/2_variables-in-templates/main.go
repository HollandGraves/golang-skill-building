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

// //////////////////////////////////////////////////////////////////////////////////
// MAIN FUNCTION
// //////////////////////////////////////////////////////////////////////////////////

func main() {
	/*
		the string in the third argument will get passed into {{.}} as before
		this time, within the template, there is a variable set up to recieve
		the . (i.e. data of the third argument)
		{{$wisdom := .}} is the variable assigment structure
		it uses a $ which is similar to PHP variable declaration
		as well as the := short variable declaration of GOlang
	*/
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", `life with a meaning`)
	if err != nil {
		log.Fatalln(err)
	}
}
