package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatal("error parsing files", err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal("error during template execution", err)
	}

	// practicing creating a template type to contain all the templates,
	// then executing the template to the stdout
	// NOT STANDARD PRACTICE CODE FOR ACTUAL PROJECT
}
