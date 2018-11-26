package main

/*
- Create a new type: vehicle with underlying type is a struct
- The fields: doors, color
- Create two new types: truck & sedan with both underlying types as struct
- Embed the “vehicle” type in both truck & sedan
- Give truck the field “fourWheel” which will be set to bool
- Give sedan the field “luxury” which will be set to bool.
- Using the vehicle, truck, and sedan structs: using a composite literal create a value of type truck and assign values to the fields
- using a composite literal, create a value of type sedan and assign values to the fields
- Print out each of these values
- Print out a single field from each of these values
- Give a method to both the “truck” and “sedan” types with the following signature transportationDevice() string
- Have each func return a string saying what they do
- Call the method for each value.
- Create a new type called “transportation” with underlying type interface.
- Create a func called “report” that takes a value of type “transportation” as an argument
- The func should print the string returned by “transportationDevice()” being called for any type that implements the “transportation” interface
- Call “report” passing in a value of type truck
- Call “report” passing in a value of type sedan.
*/

import (
	"fmt"
)

// //////////////////////////////////////////////////////////////////////////////////
// DATA DECLARATIONS
// //////////////////////////////////////////////////////////////////////////////////

type vehicle struct {
	doors string
	color string
}

type truck struct {
	fourWheel bool
	vehicle
}

type sedan struct {
	luxury bool
	vehicle
}

type transportation interface {
	transportationDevice() string
}

// //////////////////////////////////////////////////////////////////////////////////
// PACKAGE FUNCTIONS
// //////////////////////////////////////////////////////////////////////////////////

func (t truck) transportationDevice() string {
	return fmt.Sprintln("A truck is a vehicle that is used for moving heavy items")
}

func (s sedan) transportationDevice() string {
	return fmt.Sprintln("A sedan is a vehicle that is used to transport people")
}

func report(t transportation) {
	t.transportationDevice()
}

// //////////////////////////////////////////////////////////////////////////////////
// MAIN FUNCTION
// //////////////////////////////////////////////////////////////////////////////////

func main() {
	t1 := truck{
		fourWheel: true,
		vehicle: vehicle{
			doors: "4",
			color: "Blue",
		},
	}

	s1 := sedan{
		luxury: true,
		vehicle: vehicle{
			doors: "4",
			color: "Red",
		},
	}

	fmt.Println(t1)
	fmt.Println(s1)
	// a more specific way: fmt.Println(t1.vehicle.color)
	fmt.Println(t1.color)
	// a more specific way: fmt.Println(s1.vehicle.color)
	fmt.Println(s1.color)

	truckUse := t1.transportationDevice()
	sedanUse := s1.transportationDevice()
	fmt.Println(truckUse)
	fmt.Println(sedanUse)

	report(t1)
	report(s1)
}
