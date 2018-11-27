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
	sages := map[string]string{
		"Jesus":    "spoke of universal love",
		"MLK":      "spoke of acceptance of all peoples, notably African Americans",
		"Gandhi":   "spoke of Indian independence from the British, and of nonviolence",
		"Mohammad": "spoke of how God sent his final message to man through Mohammah",
		"Buddha":   "spoke of being within the moment, and nowness",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", sages)
	if err != nil {
		log.Fatalln(err)
	}
}
