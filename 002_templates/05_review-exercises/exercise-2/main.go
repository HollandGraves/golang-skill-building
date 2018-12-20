package main

/*
	Create a data structure to pass to a template which
	contains information about restaurant's menu including Breakfast, Lunch, and Dinner items
*/

import (
	"log"
	"os"
	"text/template"
)

// //////////////////////////////////////////////////////////////////////////////////
// DATA DECLARATIONS
// //////////////////////////////////////////////////////////////////////////////////

// Food : contains information about a food
type Food struct {
	Name    string
	Cal     float64
	Protein float64
	Carb    float64
	Fat     float64
}

// Restaurant : contains information about a restaurant
type Restaurant struct {
	Name      string
	Hours     string
	Breakfast []Food
	Lunch     []Food
	Dinner    []Food
}

var tpl *template.Template

// //////////////////////////////////////////////////////////////////////////////////
// PACKAGE FUNCTIONS
// //////////////////////////////////////////////////////////////////////////////////

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

// //////////////////////////////////////////////////////////////////////////////////
// MAIN FUNCTION
// //////////////////////////////////////////////////////////////////////////////////

func main() {
	/*
		ALL FOOD IS PER 100 GRAMS
	*/
	albacore := Food{Name: "Albacore Tuna", Cal: 128, Protein: 24, Carb: 0, Fat: 3}

	avocado := Food{Name: "Avocado", Cal: 167, Protein: 2, Carb: 8.6, Fat: 15}

	babySpinach := Food{Name: "Baby Spinach", Cal: 23, Protein: 2.9, Carb: 3.6, Fat: .4}

	gtKombucha := Food{Name: "GT's Gingerade Kombucha", Cal: 14, Protein: 1.2, Carb: 2.3, Fat: 0.2}

	quinoa := Food{Name: "Quinoa", Cal: 120, Protein: 4.4, Carb: 21, Fat: 1.9}

	seitan := Food{Name: "Seitan", Cal: 126, Protein: 25, Carb: 5.3, Fat: 0.6}

	sourdoughToast := Food{Name: "Sourdough Toast", Cal: 319, Protein: 13, Carb: 62, Fat: 2.1}

	steak := Food{Name: "Steak", Cal: 278, Protein: 26, Carb: 0, Fat: 18}

	/*
		RESTAURANTS
	*/
	bvc := Restaurant{
		Name:      "Basic Vegan Choices",
		Hours:     "10:00 A.M. - 8:00 P.M.",
		Breakfast: []Food{avocado, sourdoughToast, gtKombucha},
		Lunch:     []Food{avocado, babySpinach, gtKombucha, quinoa, seitan},
		Dinner:    []Food{avocado, babySpinach, gtKombucha, quinoa, seitan, sourdoughToast},
	}

	hdr := Restaurant{
		Name:      "Holland's Dream Restaurant",
		Hours:     "6:00 A.M. - 12:00 A.M.",
		Breakfast: []Food{avocado, sourdoughToast, gtKombucha},
		Lunch:     []Food{albacore, avocado, babySpinach, gtKombucha, quinoa, seitan, steak},
		Dinner:    []Food{albacore, avocado, babySpinach, gtKombucha, quinoa, seitan, sourdoughToast, steak},
	}

	/*
		ANONYMOUS STRUCT TO PASS INTO TEMPLATE AS DATA
	*/
	restaurants := []Restaurant{bvc, hdr}

	data := struct {
		GetFoodPlace []Restaurant
	}{
		restaurants,
	}

	/*
		EXECUTING DATA INTO TEMPLATE AND WRITING TO STDOUT
	*/
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
