//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

func main() {
	// 1. Declare a variable named `favoriteColor` and assign it your favorite color
	var favoriteColor = "blue"
	fmt.Print("My favorite color is ", favoriteColor, "\n")
	// 2. Declare a variable named `age` and assign it your age
	birthYear, ageInYear := 1997, 25
	fmt.Print("I was born in ", birthYear, " and I am ", ageInYear, " years old.\n")
	// 3. Declare a variable named `firstInitial` and assign it your first initial
	var (
		firstInitial = "J"
		lastInitial  = "E"
	)
	fmt.Print("My initials are ", firstInitial, lastInitial, ".\n")

	// 4. Declare a variable named `ageInDays` and assign it your age in days
	var ageInDays int
	ageInDays = 365 * ageInYear
	fmt.Print("I am ", ageInDays, " days old.\n")
}
