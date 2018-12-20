package main

import (
	"html/template"
	"log"
	"os"
)

/*
	Create a data structure to pass to a template which
	contains information about California hotels including Name, Address, City, Zip, Region
	region can be: Southern, Central, Northern
	can hold an unlimited number of hotels
*/

// //////////////////////////////////////////////////////////////////////////////////
// DATA DECLARATIONS
// //////////////////////////////////////////////////////////////////////////////////

// CaliHotels : contains information on California Hotels
type CaliHotels struct {
	Name    string
	Address string
	City    string
	Zip     int
	Region  string
}

// Hotels : contains slices of structs of California Hotels and any other hotels needed
type Hotels struct {
	CaliHotels []CaliHotels
}

// //////////////////////////////////////////////////////////////////////////////////
// PACKAGE FUNCTIONS
// //////////////////////////////////////////////////////////////////////////////////

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

// //////////////////////////////////////////////////////////////////////////////////
// MAIN FUNCTION
// //////////////////////////////////////////////////////////////////////////////////

func main() {

	// RAW DATA
	ritzLA := CaliHotels{
		Name:    "The Ritz-Carlton",
		Address: "900 W Olympic Blvd",
		City:    "Los Angeles",
		Zip:     9001,
		Region:  "Southern California",
	}
	hotelFullerton := CaliHotels{
		Name:    "The Hotel Fullerton",
		Address: "1500 S Raymond Ave",
		City:    "Fullerton",
		Zip:     92831,
		Region:  "Southern California",
	}
	appleFarm := CaliHotels{
		Name:    "Apple Farm",
		Address: "2015 Monterery St",
		City:    "San Luis Obispo",
		Zip:     93401,
		Region:  "Central California",
	}
	postRanchInn := CaliHotels{
		Name:    "Post Ranch Inn",
		Address: "47900 CA-1",
		City:    "Big Sur",
		Zip:     93920,
		Region:  "Central California",
	}
	fairmontHeritagePlace := CaliHotels{
		Name:    "Fairmont Heritage Place, Ghirardelli Square",
		Address: "900 North Point St d100",
		City:    "San Francisco",
		Zip:     94109,
		Region:  "Northern California",
	}
	alpineLodge := CaliHotels{
		Name:    "Alpine Lodge",
		Address: "908 S Mt Shasta Blvd",
		City:    "Mt Shasta",
		Zip:     96067,
		Region:  "Northern California",
	}

	// DATA FOR TEMPLATE
	Hotels := Hotels{
		CaliHotels: []CaliHotels{ritzLA, hotelFullerton, appleFarm, postRanchInn, fairmontHeritagePlace, alpineLodge},
	}

	// PASSING DATA INTO TEMPLATE
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", Hotels)
	if err != nil {
		log.Fatalln(err)
	}
}
