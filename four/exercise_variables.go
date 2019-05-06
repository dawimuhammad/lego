package main

import "fmt"

var isLiquidos bool

// ExerciseInt function contains exercise with integer variable
func ExerciseInt() {
	// ---------------------------------------------------------
	// EXERCISE: Declare int
	//
	//  1. Declare and print a variable with an int type
	//
	//  2. The declared variable's name should be: height
	//
	// EXPECTED OUTPUT
	//  0
	// ---------------------------------------------------------

	var height int

	fmt.Println("height : ", height)
	fmt.Printf("height (%T) : \n", height)

}

// ExerciseBool function contains exercise with bool variable
func ExerciseBool() {
	// ---------------------------------------------------------
	// EXERCISE: Declare bool
	//
	//  1. Declare and print a bool variable
	//
	//  2. The variable's name should be: isOn
	//
	// EXPECTED OUTPUT
	//  false
	// ---------------------------------------------------------

	var isOn bool

	fmt.Println("isOn : ", isOn)
}

// ExerciseFloat function contains exercise with float variable
func ExerciseFloat() {
	// ---------------------------------------------------------
	// EXERCISE: Declare float64
	//
	//  1. Declare and print a variable with a float64 type
	//
	//  2. The declared variable's name should be: brightness
	//
	// EXPECTED OUTPUT
	//  0
	// ---------------------------------------------------------

	var brightness float64

	fmt.Println("brightness : ", brightness)
}

// ExerciseString function contains exercise with float variable
func ExerciseString() {
	// ---------------------------------------------------------
	// EXERCISE: Declare string
	//
	//  1. Declare a string variable
	//
	//  2. Print that variable
	//
	// EXPECTED OUTPUT
	//  ""
	// ---------------------------------------------------------

	var mssg string

	fmt.Println("mssg : ", mssg)

	// USE THE BELOW CODE
	// You'll learn about Printf later

	var s string
	fmt.Printf("s (%T): %q\n", s, s)

	// %T prints the type of the value
	// %q prints an empty string
}

// ExerciseUndeclarables contains exercise with undeclarable name variables
func ExerciseUndeclarables() {
	// ---------------------------------------------------------
	// EXERCISE: Undeclarables
	//
	//  1. Declare the variables below:
	//      3speed
	//      !speed
	//      spe?ed
	//      var
	//      func
	//      package
	//
	//  2. Observe the error messages
	//
	// NOTE
	//  The types of the variables are not important.
	// ---------------------------------------------------------

	// var 3speed int
	// var !speed int
	// var spe?ed int
	// var var int
	// var func int
	// var package int

}

// ExerciseWithBits contains exercise with all variable types to it's bits
func ExerciseWithBits() {
	// ---------------------------------------------------------
	// EXERCISE: Declare with bits
	//
	//  1. Declare a few variables using the following types
	//    int
	//    int8
	//    int16
	//    int32
	//    int64
	//    float32
	//    float64
	//    complex64
	//    complex128
	//    bool
	//    string
	//    rune
	//    byte
	//
	// 2. Observe their output
	//
	// 3. After you've done, check out the solution
	//    and read the comments there
	//
	// EXPECTED OUTPUT
	//  0 0 0 0 0 0 0 (0+0i) (0+0i) false 0 0
	//  ""
	// ---------------------------------------------------------

	var i int
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64

	fmt.Println("Integer types : ", i, i8, i16, i32, i64)

	var f32 float32
	var f64 float64

	fmt.Println("Float types : ", f32, f64)

	var c64 complex64
	var c128 complex128
	var b bool
	var s string
	var r rune
	var by byte

	fmt.Println("Complex types : ", c64, c128)

	fmt.Println("Boolean type : ", b)

	fmt.Println("String type : ", s)

	fmt.Println("Rune type : ", r)

	fmt.Println("Byte type : ", by)
}

// ExerciseMultiple contains exercise with multiple variable declaration
func ExerciseMultiple() {
	// ---------------------------------------------------------
	// EXERCISE: Multiple
	//
	//  1. Declare two variables using
	//     multiple variable declaration statement
	//
	//  2. The first variable's name should be active
	//  3. The second variable's name should be delta
	//
	//  4. Print them all
	//
	// HINT
	//  You should declare a bool and an int variable
	//
	// EXPECTED OUTPUT
	//  false 0
	// ---------------------------------------------------------

	var (
		active, delta int
	)

	fmt.Println("Multiple declarations : ", active, delta)
}

// ExerciseMultipleTwo contains exercise with other multiple variable declaration
func ExerciseMultipleTwo() {
	// ---------------------------------------------------------
	// EXERCISE: Multiple #2
	//
	//  1. Declare and initialize two string variables
	//     using multiple variable declaration
	//
	//  2. Use the type once while declaring the variables
	//
	//  3. The first variable's name should be firstName
	//  4. The second variable's name should be lastName
	//
	//  5. Print them all
	//
	// EXPECTED OUTPUT
	//  "" ""
	// ---------------------------------------------------------

	var firstName, lastName string = "", ""

	firstName = "Muhamad Haddawi"
	lastName = "Hermawan"

	fmt.Println("firstName : ", firstName)
	fmt.Println("lastName : ", lastName)
}

// ExerciseUnusedVariable contains exercise to know how to use blank-identifier
func ExerciseUnusedVariable() {
	// ---------------------------------------------------------
	// EXERCISE: Unused
	//
	//  1. Declare a variable
	//
	//  2. Variable's name should be: isLiquid
	//
	//  3. Discard it using a blank-identifier
	//
	// NOTE
	//  Do not print the variable
	// ---------------------------------------------------------

	var isLiquid string
	_ = isLiquid
}

// ExerciseWrongDoer contains exercise to observe the error on declaring variable not on the first
func ExerciseWrongDoer() {
	// ---------------------------------------------------------
	// EXERCISE: Wrong doer
	//
	//  1. Print a variable
	//
	//  2. Then declare it
	//  (This means: Try to print it before its declaration)
	//
	//  3. Observe the error
	// ---------------------------------------------------------

	// fmt.Println("Wrong doer", doer)

	// var doer string
}
