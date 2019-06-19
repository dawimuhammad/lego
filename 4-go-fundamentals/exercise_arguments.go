package main

import (
	"fmt"
	"os"
)

// CountArgs contains exercise on course arguments to print the count of given arguments
func CountArgs() {
	// ---------------------------------------------------------
	// EXERCISE: Count the Arguments
	//
	//  Print the count of the command-line arguments
	//
	// INPUT
	//  bilbo balbo bungo
	//
	// EXPECTED OUTPUT
	//  There are 3 names.
	// ---------------------------------------------------------

	// pseudo style

	// var count int

	// count = len(os.Args) - 1

	// fmt.Println("\nCountArgs()")
	// fmt.Println("Args length : ", count)

	// short code style

	count := len(os.Args) - 1

	fmt.Println("Args Len : ", count)
}

// PrintPath contains exercise on course arguments to print the command path
func PrintPath() {
	// ---------------------------------------------------------
	// EXERCISE: Print the Path
	//
	//  Print the path of the running program by getting it
	//  from `os.Args` variable.
	//
	// HINT
	//  Use `go build` to build your program.
	//  Then run it using the compiled executable program file.
	//
	// EXPECTED OUTPUT SHOULD INCLUDE THIS
	//  myprogram
	// ---------------------------------------------------------

	path := os.Args[0]

	fmt.Println("PrintPath()\nIt Prints : ", path)
}

// PrintNames returns given names from arguments
func PrintNames() {
	// ---------------------------------------------------------
	// EXERCISE: Greet 5 People
	//
	//  Greet 5 people this time.
	//
	//  Please do not copy paste from the previous exercise!
	//
	// RESTRICTION
	//  This time do not use variables.
	//
	// INPUT
	//  bilbo balbo bungo gandalf legolas
	//
	// EXPECTED OUTPUT
	//  There are 5 people!
	//  Hello great bilbo !
	//  Hello great balbo !
	//  Hello great bungo !
	//  Hello great gandalf !
	//  Hello great legolas !
	//  Nice to meet you all.
	// ---------------------------------------------------------

	nameLength := len(os.Args)

	for i := 1; i < nameLength; i++ {
		fmt.Println("PrintNames()\nOn index", i, "it prints : ", os.Args[i])
	}
}
