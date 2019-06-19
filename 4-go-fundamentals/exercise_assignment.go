package main

import (
	"fmt"
	"path"
)

// MakeItBlue contains exercise to change variable value using assignment
func MakeItBlue() {
	// ---------------------------------------------------------
	// EXERCISE: Make It Blue
	//
	//  1. Change `color` variable's value to "blue"
	//
	//  2. Print it
	//
	// EXPECTED OUTPUT
	//  blue
	// ---------------------------------------------------------

	fmt.Println("\nMakeItBlue()")

	var color = "red"

	fmt.Println("Prev Color : ", color)

	color = "blue"

	fmt.Println("Current Color : ", color)
}

// VarToVar contains exercise to change value using variables swap
func VarToVar() {
	// ---------------------------------------------------------
	// EXERCISE: Variables To Variables
	//
	//  1. Change the value of `color` variable to "dark green"
	//
	//  2. Do not assign "dark green" to `color` directly
	//
	//     Instead, use the `color` variable again
	//     while assigning to `color`
	//
	//  3. Print it
	//
	// RESTRICTIONS
	//  WRONG ANSWER, DO NOT DO THIS:
	//  `color = "dark green"`
	//
	// HINT
	//  + operator can concatenate string values
	//
	//  Example:
	//
	//  "h" + "e" + "y" returns "hey"
	//
	// EXPECTED OUTPUT
	//  dark green
	// ---------------------------------------------------------

	var color string

	color = "dark" + " " + "green"

	fmt.Println("\nVarToVar()")

	fmt.Println("Color : ", color)
}

// AssignWithExpression contains exercise to learn assign variable with combination of expression
func AssignWithExpression() {
	// ---------------------------------------------------------
	// EXERCISE: Assign With Expressions
	//
	//  1. Multiply 3.14 with 2 and assign it to `n` variable
	//
	//  2. Print the `n` variable
	//
	// HINT
	//  Example: 3 * 2 = 6
	//  * is the product operator (it multiplies numbers)
	//
	// EXPECTED OUTPUT
	//  6.28
	// ---------------------------------------------------------

	fmt.Println("\nAssignWithExpression()")

	var n float64

	n = 3.14 * 2

	fmt.Println("n value is ", n)
}

// FindPerimeter contains exercise to learn assignment with expressions
func FindPerimeter() {
	// ---------------------------------------------------------
	// EXERCISE: Find the Rectangle's Perimeter
	//
	//  1. Find the perimeter of a rectangle
	//     Its width is  5
	//     Its height is 6
	//
	//  2. Assign the result to the `perimeter` variable
	//
	//  3. Print the `perimeter` variable
	//
	// HINT
	//  Formula = 2 times the width and height
	//
	// EXPECTED OUTPUT
	//  22
	//
	// BONUS
	//  Find more formulas here and calculate them in new programs
	//  https://www.mathsisfun.com/area.html
	// ---------------------------------------------------------

	var (
		perimeter     int
		width, height = 5, 6
	)

	perimeter = (width * 2) + (height * 2)
	fmt.Println("\nFindPerimeter()")

	fmt.Println("Perimeter : ", perimeter)
}

// MultiAssignment contains exercise to learn multi assignment vol 1
func MultiAssignment() {
	// ---------------------------------------------------------
	// EXERCISE: Multi Assign
	//
	//  1. Assign "go" to `lang` variable
	//     and assign 2 to `version` variable
	//     using a multiple assignment statement
	//
	//  2. Print the variables
	//
	// EXPECTED OUTPUT
	//  go version 2
	// ---------------------------------------------------------

	var (
		lang    string
		version int
	)

	lang, version = "go", 2

	fmt.Println("\nMultiAssignment1()")

	fmt.Println(lang, " ", version)

	var (
		planet string
		isTrue bool
		temp   float64
	)

	planet, isTrue, temp = "Air is good on Mars", true, 19.5

	fmt.Println("\nMultiAssignment2()")

	fmt.Println("Planet desc : ", planet)
	fmt.Println("Can we live here : ", isTrue)
	fmt.Println("Planet desc : ", temp)

}

// MultiShortFunc contains exercise to learn multiple short declaration (again)
func MultiShortFunc() {
	// ---------------------------------------------------------
	// EXERCISE: Multi Short Func
	//
	// 	1. Declare two variables using multiple short declaration syntax
	//
	//  2. Initialize the variables using `multi` function below
	//
	//  3. Discard the 1st variable's value in the declaration
	//
	//  4. Print only the 2nd variable
	//
	// NOTE
	//  You should use `multi` function
	//  while initializing the variables
	//
	// EXPECTED OUTPUT
	//  4
	// ---------------------------------------------------------

	fmt.Println("\nMultiShortFunc()")

	// variables declarations
	var a, b int

	// initialize variables using multi func
	a, b = multi()

	fmt.Println("Initial Value of a : ", a, "; b : ", b)

	// discard first variable
	_ = a

	fmt.Println("Print b : ", b)

	// ---------  SHOT CODE

	_, otherB := multi()

	fmt.Println("other b : ", otherB)
}

func multi() (int, int) {
	return 5, 4
}

// SwapperFirst contains exercise to learn multi-assignment in go
func SwapperFirst() {
	// ---------------------------------------------------------
	// EXERCISE: Swapper
	//
	//  1. Change `color` to "orange"
	//     and `color2` to "green" at the same time
	//
	//     (use multiple-assignment)
	//
	//  2. Print the variables
	//
	// EXPECTED OUTPUT
	//  orange green
	// ---------------------------------------------------------

	fmt.Println("\nSwapperFirst()")

	color, color2 := "orange", "green"

	fmt.Println("Initial values\ncolor : ", color, "\ncolor2 : ", color2)

	color, color2 = color2, color

	fmt.Println("After values\ncolor : ", color, "\ncolor2 : ", color2)
}

// DiscardFile returns directory of a given file pathname
func DiscardFile() {
	// ---------------------------------------------------------
	// EXERCISE: Discard The File
	//
	//  1. Print only the directory using `path.Split`
	//
	//  2. Discard the file part
	//
	// RESTRICTION
	//  Use short declaration
	//
	// EXPECTED OUTPUT
	//  secret/
	// ---------------------------------------------------------

	dir, _ := path.Split("secret/key.txt")

	fmt.Println("DiscardFile()")
	fmt.Println("Directory path : ", dir)
}
