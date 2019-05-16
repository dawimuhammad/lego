package main

import "fmt"

// ConvertAndFixOne contains exercise convert and fix part one
func ConvertAndFixOne() {
	// ---------------------------------------------------------
	// EXERCISE: Convert and Fix
	//
	//  Fix the code by using the conversion expression.
	//
	// EXPECTED OUTPUT
	//  15.5
	// ---------------------------------------------------------

	a, b := 10, 5.5
	fmt.Println(float64(a) + b)

}

// ConvertAndFixTwo contains exercise convert and fix part two
func ConvertAndFixTwo() {
	// ---------------------------------------------------------
	// EXERCISE: Convert and Fix #2
	//
	//  Fix the code by using the conversion expression.
	//
	// EXPECTED OUTPUT
	//  10.5
	// ---------------------------------------------------------

	a, b := 10, 5.5
	a = int(b)
	fmt.Println(float64(a) + b)

}

// ConvertAndFixThree contains exercise convert and fix part three
func ConvertAndFixThree() {
	// ---------------------------------------------------------
	// EXERCISE: Convert and Fix #4
	//
	//  Fix the code.
	//
	// EXPECTED OUTPUT
	//  9.5
	// ---------------------------------------------------------

	age := 2

	fmt.Println((7.5) + float64(age))

}

// ConvertAndFixFour contains exercise convert and fix part Four
func ConvertAndFixFour() {
	// ---------------------------------------------------------
	// EXERCISE: Convert and Fix #5
	//
	//  Fix the code.
	//
	// HINTS
	//   maximum of int8  can be 127
	//   maximum of int16 can be 32767
	//
	// EXPECTED OUTPUT
	//  1127
	// ---------------------------------------------------------

	// DO NOT TOUCH THIS VARIABLES
	min := int8(127)
	max := int16(1000)

	// FIX THE CODE HERE
	fmt.Println(int16(max) + int16(min))

}
