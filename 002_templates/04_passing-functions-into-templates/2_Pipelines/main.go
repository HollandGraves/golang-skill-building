package main

/*
	1. pass functions into template
	2. use functions in template and pipline the result into other functions
*/

import (
	"html/template"
	"log"
	"math"
	"os"
)

// //////////////////////////////////////////////////////////////////////////////////
// DATA DECLARATIONS
// //////////////////////////////////////////////////////////////////////////////////

// //////////////////////////////////////////////////////////////////////////////////
// PACKAGE FUNCTIONS
// //////////////////////////////////////////////////////////////////////////////////

var fm = template.FuncMap{
	"di":   DoubleInt,
	"sn":   SquareNum,
	"nitf": NumIntToFloat64,
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

// DoubleInt : doubles int lol
func DoubleInt(x int) int {
	return x * 2
}

// NumIntToFloat64 : turns an int into float64 type
func NumIntToFloat64(x int) float64 {
	return float64(x)
}

// SquareNum : well, it squares num!
func SquareNum(x float64, y float64) float64 {
	return math.Pow(x, y)
}

// //////////////////////////////////////////////////////////////////////////////////
// MAIN FUNCTION
// //////////////////////////////////////////////////////////////////////////////////

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
