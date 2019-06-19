package main

import (
	"fmt"
	"os"
)

// ArgsExercise returns input arguments
func ArgsExercise() {
	// var Args []string

	fmt.Printf("%#v\n", os.Args)

	argsLength := len(os.Args)

	for i := 0; i < argsLength; i++ {
		fmt.Println("slice ", i, " contains : ", os.Args[i])
	}
}

// GreetPeople return greeting into given peoples name
func GreetPeople() {
	peoples := len(os.Args)

	for i := 1; i < peoples; i++ {
		people := os.Args[i]
		fmt.Println("Hello ", people, ", have a nice day!")
	}
}
