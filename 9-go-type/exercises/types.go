package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

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

// ParseArgNumbers contain go exercise to learn how to parse arg numbers
func ParseArgNumbers() {
	// ---------------------------------------------------------
	// EXERCISE: Parse Arg Numbers
	//
	//  Use strconv.ParseInt function to get int8, int16, and
	//  int32, and int64 values from command-line.
	//
	// HINT
	//  The third argument to ParseInt function represents
	//  the bitsize.
	//
	//  So, giving it 8 returns an int8 convertable value;
	//  whereas 16 returns an int16 convertable value.
	//
	//  Please explore the documentation of ParseInt function
	//  and learn how it works.
	//
	// EXPECTED OUTPUT
	//   When runned like this:
	//     go run main.go 50 25000 2000000 50000000000000000 00000100
	//
	//   It should return this:
	//     int8 value is : 50
	//     int16 value is: 25000
	//     int32 value is: 2000000
	//     int64 value is: 50000000000000000
	//     00000100 is: 4
	// ---------------------------------------------------------

	// --------------------------------------
	// EXAMPLE:
	// --------------------------------------
	// How to get an int8 from command-line:
	// First argument should be a value of -128 to 127 range
	//
	// Second argument: 10 means decimal number
	// Third argument : 8 means 8-bits (int8)
	val, _ := strconv.ParseInt(os.Args[1], 10, 8)

	// Now the val variable is int64 because ParseInt
	// returns an int64. But, since I passed 8 as its third
	// argument, it returns int8 convertable value.
	//
	// Try running the program with a value of -128 to 127
	// Running it beyond that range will result in
	// either -128 or 127.
	fmt.Printf("%T value is : %d \n", int8(val), int8(val))

	// --------------------------------------
	// NOW IT'S YOUR TURN!
	// --------------------------------------

	// 1. Get an int16 value using ParseInt and convert it and print it
	val, _ = strconv.ParseInt(os.Args[2], 10, 16)
	fmt.Printf("%T value is : %d \n", int16(val), int16(val))

	// 2. Get an int32 value using ParseInt and convert it and print it
	val, _ = strconv.ParseInt(os.Args[3], 10, 32)
	fmt.Printf("%T value is : %d \n", int32(val), int32(val))

	// 3. Get an int64 value using ParseInt and convert it and print it
	val, _ = strconv.ParseInt(os.Args[4], 10, 64)
	fmt.Printf("%T value is : %d \n", int64(val), int64(val))

	// 4. Get an int8 value using ParseInt and convert it and print it
	//    But this time, get the value in bits.
	//
	//    For example : 00000100
	//    Should print: 4

	val, _ = strconv.ParseInt(os.Args[5], 2, 8)

	fmt.Printf("%s is : %d \n", os.Args[5], int8(val))
}

// TimeMultiplier contains exercise to converts input value to time value
func TimeMultiplier() {
	// ---------------------------------------------------------
	// EXERCISE: Time Multiplier
	//
	//  You should get an argument from the command-line and
	//  you need to multiply the time duration value `t` with
	//  the given argument.
	//
	//  1. Get an argument from the command-line
	//  2. Convert it to int64 and store it in a variable
	//  3. Multiply the `t` variable with that variable
	//  4. Print the result
	//
	// HINT
	//  You can use ParseInt to convert the command-line
	//    argument into an int64 value.
	//
	//  You can skip the error value using a blank-identifier.
	//
	// EXPECTED OUTPUT
	//
	//  When runned like this:
	//    go run main.go 2
	//
	//  It should print this:
	//    3h0m0s
	// ---------------------------------------------------------

	// DONT TOUCH THIS
	// 1h30m means: 1 hour 30 minutes
	t, _ := time.ParseDuration("1h30m")

	// TYPE YOUR CODE HERE
	// ....
	multiplier, _ := strconv.ParseInt(os.Args[1], 10, 64)
	t *= time.Duration(multiplier)

	// DONT TOUCH THIS
	fmt.Println(t)
}

// RefactorFeetToMeter contains exercise to learn underlying types with previously written code
func RefactorFeetToMeter() {
	// ---------------------------------------------------------
	// EXERCISE: Refactor Feet to Meter
	//
	//  Remember the feet to meters program?
	//  Now, it's time to refactor it.
	//  Define your own Feet and Meters types.
	//
	//  Follow the steps inside the code.
	// ---------------------------------------------------------

	// ----------------------------
	// 1. Define Feet and Meters types below
	//    Their underlying type can be float64

	type Feet float64
	type Meters float64

	// ----------------------------
	// 2. UNCOMMENT THE CODE BELOW THEN DON'T TOUCH IT

	var (
		feet   Feet
		meters Meters
	)

	// ----------------------------
	// 3. Get feet value from the command-line
	// 4. Convert it to an float64 first using ParseFloat
	// 5. Then, convert it into a Feet type

	feetFloat, _ := strconv.ParseFloat(os.Args[1], 64)

	feet = Feet(feetFloat)

	// ----------------------------
	// 6. Uncomment the code below
	// 7. And, convert the expression to Meters type

	meters = Meters(feet) * 0.3048

	// ----------------------------
	// 8. UNCOMMENT THE CODE BELOW
	fmt.Printf("%g feet is %g meters.\n", feet, meters)
}
