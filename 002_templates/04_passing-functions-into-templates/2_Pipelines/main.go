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

var fm = template.FuncMap{
	"di": DoubleInt,
	"sn": SquareNum,
	// "nitf": NumIntToFloat64, this function wont work for reasons below
}

var tpl *template.Template

// //////////////////////////////////////////////////////////////////////////////////
// PACKAGE FUNCTIONS
// //////////////////////////////////////////////////////////////////////////////////

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

// DoubleInt : doubles int lol
func DoubleInt(x int) int {
	return x * 2
}

// This function will not be able to be passed because it
// returns two values and one of the values types is not of
// type err.
// The only functions that can compile into templates are
// functions that have one return value, or two return values
// with one being of type err and the other being of another type
//
// func NumIntToFloat64(x int, y int) (float64, float64) {
// 	return float64(x), float64(y)
// }

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
