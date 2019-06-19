package main

import "fmt"

// DoSomeCalc contains some calculation exercise
func DoSomeCalc() {
	// ---------------------------------------------------------
	// EXERCISE: Do Some Calculations
	//
	//  1. Print the sum of 50 and 25
	//  2. Print the difference of 50 and 15.5
	//  3. Print the product of 50 and 0.5
	//  4. Print the quotient of 50 and 0.5
	//  5. Print the remainder of 25 and 3
	//  6. Print the negation of `5 + 2`
	//
	// EXPECTED OUTPUT
	//  75
	//  34.5
	//  25
	//  100
	//  1
	//  -7
	// ---------------------------------------------------------

	fmt.Printf("the sum of 50 and 25 is %d \n", 50+25)

	fmt.Printf("the substract of 50 with 15.5 is %g \n", 50-15.5)

	fmt.Printf("the product between 50 and 0.5 is %g\n", 50*.5)

	fmt.Printf("the quotient of 50 by 0.5 is %g\n", 50/0.5)

	fmt.Printf("the remainder of 25 by 3 is %d\n", 25%3)

	fmt.Printf("the negation from the result of 5 + 2 is %d\n", -(5 + 2))
}

// FixTheFloat contains exercise to fix the wrong float calculation
func FixTheFloat() {
	// ---------------------------------------------------------
	// EXERCISE: Fix the Float
	//
	//  Fix the program to print 2.5 instead of 2
	//
	// EXPECTED OUTPUT
	//  2.5
	// ---------------------------------------------------------

	x := float64(5) / 2
	fmt.Println("Fix The Float : ", x)
}

// TryPrecedence contains exercise to learn and practice of using precedence
func TryPrecedence() {
	// ---------------------------------------------------------
	// EXERCISE: Precedence
	//
	//  Change the expressions to produce the expected outputs
	//
	// RESTRICTION
	//  Use parentheses to change the precedence
	// ---------------------------------------------------------

	fmt.Printf("\nTryPrecedence()\n")

	// This expression should print 20
	fmt.Println("20 prints ", 10+5-(5-10))

	// This expression should print -16
	fmt.Println("-16 prints ", -10+0.5-(1+5.5))

	// This expression should print -25
	fmt.Println("-25 prints ", 5+10*(2-5))

	// This expression should print 0.5
	fmt.Println("0.5 prints ", 0.5*(2-1))

	// This expression should print 24
	fmt.Println("24 prints ", ((3+1)/2)*10+4)

	// This expression should print 15
	fmt.Println("15 prints ", (10/2)*(10%7))

	// This expression should print 40
	// Note that you may need to use floats to solve this
	fmt.Println("40 should prints ", 100/(float64(5)/2))
}

// Indec func contains exercise to learn and practice indec
func Indec() {
	// ---------------------------------------------------------
	// EXERCISE: Incdecs
	//
	//  1. Increase the `counter` 5 times
	//  2. Decrease the `factor` 2 times
	//  3. Print the product of counter and factor
	//
	// RESTRICTION
	//  Use only the incdec statements
	//
	// EXPECTED OUTPUT
	//  -75
	// ---------------------------------------------------------

	// DO NOT TOUCH THIS
	counter, factor := 45, 0.5

	// TYPE YOUR CODE BELOW
	for i := 1; i <= 5; i++ {
		counter++

		if i <= 2 {
			factor--
		}
	}

	fmt.Printf("Indec()\n")

	fmt.Printf("the result of %d multiply by %g is %g \n", counter, factor, float64(counter)*factor)

	// LASTLY: REMOVE THE CODE BELOW
	// _, _ = counter, factor
}

// ManipulateCounter contains exercise to learn to manipulate a counter
func ManipulateCounter() {
	// ---------------------------------------------------------
	// EXERCISE: Manipulate a Counter
	//
	//  1. Write the simplest line of code to increase
	//     the counter variable by 1.
	//
	//  2. Write the simplest line of code to decrease
	//     the counter variable by 1.
	//
	//  3. Write the simplest line of code to increase
	//     the counter variable by 5.
	//
	//  4. Write the simplest line of code to multiply
	//     the counter variable by 10.
	//
	//  5. Write the simplest line of code to divide
	//     the counter variable by 5.
	//
	// EXPECTED OUTPUT
	//  10
	// ---------------------------------------------------------

	// DO NOT CHANGE THE CODE BELOW
	var counter int

	// TYPE YOUR CODE HERE
	// 1.
	counter++

	// 2.
	counter--

	// 3.
	counter += 5

	// 4.
	counter *= 10

	// 5.
	counter /= 5

	fmt.Printf("Manipulate Counter result is %d \n", counter)
}

// SimplifyAssignment contains exercise indec and assignment
func SimplifyAssignment() {
	// ---------------------------------------------------------
	// EXERCISE: Simplify the Assignments
	//
	//  Simplify the code (refactor)
	//
	// RESTRICTION
	//  Use only the incdec and assignment operations
	//
	// EXPECTED OUTPUT
	//  3
	// ---------------------------------------------------------

	width, height := 10, 2

	width++
	width += height
	width--
	width -= height
	width *= 20
	width /= 25
	width %= 5

	fmt.Printf("Simplify Assignment result is %d \n", width)
}
