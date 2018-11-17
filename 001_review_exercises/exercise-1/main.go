package main

/*
- create a type square
- create a type circle
- attach a method to each that calculates area and returns it
- create a type shape which defines an interface as anything which has the area method
- create a func info which takes type shape and then prints the area
- create a value of type square
- create a value of type circle
- use func info to print the area of square
- use func info to print the area of circle
*/

import (
	"fmt"
	"math"
)

// //////////////////////////////////////////////////////////////////////////////////
// DATA DECLARATIONS
// //////////////////////////////////////////////////////////////////////////////////

// types of geometric shapes
type square struct {
	length float64
	width  float64
}
type circle struct {
	radius float64
}
type shape interface {
	findArea() float64
}

// Pi declaration
var Pi = 3.14159265358979323846264338327950288419716939937510582097494459

// //////////////////////////////////////////////////////////////////////////////////
// PACKAGE FUNCTIONS
// //////////////////////////////////////////////////////////////////////////////////

// functions to find area
func (sq *square) findArea() float64 {
	area := sq.length * sq.width
	return area
}

func (cr *circle) findArea() float64 {
	area := cr.radius * math.Pow(Pi, 2)
	return area
}

func printArea(sh shape) {
	fmt.Println(sh.findArea())
}

// //////////////////////////////////////////////////////////////////////////////////
// MAIN FUNCTION
// //////////////////////////////////////////////////////////////////////////////////

func main() {
	// gives values to the strucs
	square1 := &square{length: 8.4, width: 2}
	circle1 := &circle{radius: 4.4}

	printArea(square1)
	printArea(circle1)
}
