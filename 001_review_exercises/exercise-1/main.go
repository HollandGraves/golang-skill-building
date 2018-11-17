package main

import (
	"fmt"
	"math"
)

// types of geometric shapes
type square struct {
	length float64
	width  float64
}
type circle struct {
	radius float64
}

// Pi declarations
var Pi = 3.14159265358979323846264338327950288419716939937510582097494459

// functions to find area
func (sq square) findArea() float64 {
	area := sq.length * sq.width
	return area
}

func (cr circle) findArea() float64 {
	area := cr.radius * math.Pow(Pi, 2)
	return area
}

func main() {
	// gives values to the strucs
	square1 := square{length: 8.4, width: 2}
	circle1 := circle{radius: 4.4}

	// calculates and prints the area
	fmt.Println(square1.findArea())
	fmt.Println(circle1.findArea())
}
