package main

import (
	"fmt"
	"os"
)

// PrintingExercise returns print exercise
func PrintingExercise() {
	var (
		brand    = "Macbook Pro"
		quantity = 20
		new      = true
		price    = 21750500.75
	)

	fmt.Printf("The new notbook named %s with new condition %t is sell by price %.2f with minimum order %d pcs\n", brand, new, price, quantity)
}

// YourAges prints your current age
func YourAges() {
	// ---------------------------------------------------------
	// EXERCISE: Print Your Age
	//
	//  Print your age using Printf
	//
	// EXPECTED OUTPUT
	//  I'm 30 years old.
	//
	// NOTE
	//  You should change 30 to your age, of course.
	// ---------------------------------------------------------

	var myAge = 24

	fmt.Printf("I am %d years old.\n", myAge)
}

// YourName prints your first and your last name
func YourName() {
	// ---------------------------------------------------------
	// EXERCISE: Print Your Name and LastName
	//
	//  Print your name and lastname using Printf
	//
	// EXPECTED OUTPUT
	//  My name is Inanc and my lastname is Gumus.
	//
	// BONUS
	//  Store the formatting specifier (first argument of Printf)
	//    in a variable.
	//  Then pass it to printf
	// ---------------------------------------------------------

	var firstName, lastName = "Muhamad", "Haddawi"

	fmt.Printf("My name is %s and my lastname is %s.\n", firstName, lastName)
}

// FalseClaims prints bool on printf
func FalseClaims() {
	// ---------------------------------------------------------
	// EXERCISE: False Claims
	//
	//  Use printf to print the expected output using a variable.
	//
	// EXPECTED OUTPUT
	//  These are false claims.
	// ---------------------------------------------------------

	// UNCOMMENT THE FOLLOWING CODE
	// AND DO NOT CHANGE IT AFTERWARDS
	tf := false

	// TYPE YOUR CODE HERE
	// ?

	fmt.Printf("These are %t claims.\n", tf)
}

// WhatsTheHeat prints the temperature
func WhatsTheHeat() {
	// ---------------------------------------------------------
	// EXERCISE: Print the Temperature
	//
	//  Print the current temperature in your area using Printf
	//
	// NOTE
	//  Do not use %v verb
	//  Output "shouldn't" be like 29.500000
	//
	// EXPECTED OUTPUT
	//  Temperature is 29.5 degrees.
	// ---------------------------------------------------------

	var temperature = 29.5

	fmt.Printf("Temperature is %.1f degrees.\n", temperature)
}

// SayHello prints double quote symbol using printf
func SayHello() {
	// ---------------------------------------------------------
	// EXERCISE: Double Quotes
	//
	//  Print "hello world" with double-quotes using Printf
	//  (As you see here)
	//
	// NOTE
	//  Output "shouldn't" be just: hello world
	//
	// EXPECTED OUTPUT
	//  "hello world"
	// ---------------------------------------------------------

	fmt.Printf("\"Hello World!\"\n")
}

// CallTheType prints the value with its value-type
func CallTheType() {
	// ---------------------------------------------------------
	// EXERCISE: Print the Type
	//
	//  Print the type and value of 3 using fmt.Printf
	//
	// EXPECTED OUTPUT
	//  Type of 3 is int
	// ---------------------------------------------------------

	var qty = 3

	fmt.Printf("Type of %d is %[1]T\n", qty)
}

// PrintYourName prints your full name from given command-line input
func PrintYourName() {
	// ---------------------------------------------------------
	// EXERCISE: Print Your Fullname
	//
	//  1. Get your name and lastname from the command-line
	//  2. Print them using Printf
	//
	// EXAMPLE INPUT
	//  Inanc Gumus
	//
	// EXPECTED OUTPUT
	//  Your name is Inanc and your lastname is Gumus.
	// ---------------------------------------------------------

	firstName, lastName := os.Args[1], os.Args[2]

	fmt.Printf("Your name is %s and your lastname is %s", firstName, lastName)
}
