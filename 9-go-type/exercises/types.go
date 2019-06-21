package main

import "fmt"

// OptimalTypes contains exercise to learn many go data types
func OptimalTypes() {
	// ---------------------------------------------------------
	// EXERCISE: Optimal Types
	//
	//  1. Choose the optimal data types for the given situations.
	//  2. Print them all
	//  3. Try converting them to lesser data types.
	//     For example try converting int64 variable to int32.
	//     Then observe the result.
	//     Search the web why the result is so?
	//
	// NOTE
	//  This is just an exercise for teaching you different
	//  data types. Do not apply it to the real-life.
	//
	//  As I said in the lectures that, premature optimization
	//  is not a good thing.
	// ---------------------------------------------------------

	// DONT FORGET: There are also unsigned data types.
	//              (For positive numbers)

	// DO NOT USE: int data type
	// Use only the ones with the bitsizes

	// ---

	// an english letter (search web for: ascii control code)
	var letter byte = 'A' // equals uint8, consist of 8 bits of memory

	// a non-english letter (search web for: unicode codepoint)
	var unicodeLetter rune // rune equals int32, consist of 32 bits of memory
	unicodeLetter = 'ç' * 'ç'

	// a year in 4 digits like 2040
	var year uint16 = 2010

	// a month in 2 digits: 1 to 12
	var month uint8 = 12

	// the speed of the light
	var speedLight rune = 299792458

	// angle of a circle
	var angle float32 = 90.0

	// PI number: 3.141592653589793
	var pi float64
	pi = 3.141592653589793

	fmt.Println("letter : ", letter)
	fmt.Printf("byte : %b\n", letter)
	fmt.Println("unicodeLetter : ", unicodeLetter)
	fmt.Printf("rune : %b\n", unicodeLetter)
	fmt.Println("current year : ", year)
	fmt.Println("current month : ", month)
	fmt.Println("angle circle : ", angle)
	fmt.Println("speedlight : ", speedLight)
	fmt.Println("pi : ", pi)
}

// TypeProblem contains exercise to debug and fix the error from missmatch type
func TypeProblem() {
	// ---------------------------------------------------------
	// EXERCISE: The Type Problem
	//
	//  Solve the data type problem in the program.
	//
	// EXPECTED OUTPUT
	//  width: 265 height: 265
	//  are they equal? true
	// ---------------------------------------------------------

	// FIX THIS:
	// Change the following data types to the correct
	// data types where appropriate.

	var (
		width  uint8
		height uint16
	)

	// DONT TOUCH THIS:
	width, height = 255, 265
	width += 10
	fmt.Printf("width: %d height: %d\n", width, height)

	// UNCOMMENT THIS:
	fmt.Println("are they equal?", width == uint8(height))
}
