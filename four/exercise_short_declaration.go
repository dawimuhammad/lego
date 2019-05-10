package main

import "fmt"

// ShortDeclaration contains exercise to learn how to declares variables using short declaration in go
func ShortDeclaration() {
	// ---------------------------------------------------------
	// EXERCISE: Short Declare
	//
	//  Declare and then print four variables using
	//  the short declaration statement.
	//
	// EXPECTED OUTPUT
	//  i: 314 f: 3.14 s: Hello b: true
	// ---------------------------------------------------------

	i, f, s, b := 314, 3.14, "Hello", true

	fmt.Println("Short Declaration \n  ",
		"i : ", i,
		"f : ", f,
		"s : ", s,
		"b : ", b)
}

// ShortWithExpression contains exercise to learn how to using short variable declaration with an expression
func ShortWithExpression() {
	// ---------------------------------------------------------
	// EXERCISE: Short With Expression
	//
	// 	1. Short declare a variable named `sum`
	//
	//  2. Initialize it with an expression by adding 27 and 3.5
	//
	// EXPECTED OUTPUT
	//  30.5
	// ---------------------------------------------------------

	sum := 27 + 3.5

	fmt.Println("Short with Expression \n  ",
		"sum : ", sum)
}

// ShortDiscard contains exercise to learn how to use short declaration and discard some variable that had been declared
func ShortDiscard() {
	// ---------------------------------------------------------
	// EXERCISE: Short Discard
	//
	// 	1. Short declare two bool variables
	//     (use multiple short declaration syntax)
	//
	//  2. Initialize both variables to true
	//
	//  3. Change your declaration and
	//     discard the 2nd variable's value
	//     using the blank-identifier
	//
	//  4. Print only the 1st variable
	//
	// EXPECTED OUTPUT
	//  true
	// ---------------------------------------------------------

	first, _ := true, true

	// _ = second

	fmt.Println("Short Discard \n  ",
		"first : ", first)
}

// Redeclare contains exercise to learn how to redeclare variables that already declared before
func Redeclare() {
	// ---------------------------------------------------------
	// EXERCISE: Redeclare
	//
	// 	1. Short declare two int variables: age and yourAge
	//     (use multiple short declaration syntax)
	//
	//  2. Short declare a new float variable: ratio
	//     And, change the 'age' variable to 42
	//
	//     (! You should use redeclaration)
	//
	//  4. Print all the variables
	//
	// EXPECTED OUTPUT
	//  42, 20, 3.14
	// ---------------------------------------------------------

	age, yourAge := 20, 24

	ratio, age := 0, 42

	fmt.Println("Redeclare \n  ",
		"age : ", age,
		"yourAge : ", yourAge,
		"ratio : ", ratio)
}
