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
	var favoriteColor string = "Blue"

	fmt.Println("my favorite color is", favoriteColor)

	//	Store your birth year and age (in years) in two variables using
	//  compound assignment
	birthYear, ageInYears := 1987, 37
	fmt.Println("Boun in", birthYear, "aged", ageInYears, "years")

	//* Store your first & last initials in two variables using block assignment
	var (
		firstIntial = "M"
		lastInitial = "O"
	)
	fmt.Println("Initials=", firstIntial, lastInitial)

	//* Declare (but don't assign!) a variable for your age (in days),
	//  then assign it on the next line by multiplying 365 with the age
	// 	variable created earlier
	var ageInDays int
	ageInDays = 365 * ageInYears
	fmt.Println("I am", ageInDays, "days old")
}
