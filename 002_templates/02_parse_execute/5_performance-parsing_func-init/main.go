package main

import (
	"log"
	"os"
	"text/template"
)

// package level declaration so that any package can access the templates
var tpl *template.Template

// init() so as to get all the template to be parse before the program runs
func init() {
	// the tpl variable will act as a container for all templates
	// use of Must() so that if an error arrises the program wont run
	// use of ParseGlob() so that all templates can be parsed at once
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	/*
		- ExecuteTemplate() will Write a specific template to a location with
		data being passed in
		- In this case, the template is being written to Stdout,
		the template in question is "one.gohtml", and no data is being
		passed into the template due to the nil datatype used as an argument
	*/
	err := tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "vespa.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
