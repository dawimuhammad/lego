package main

import "fmt"

// a defined type and its source type share the same underlying type
// which is string type on the case below
type name string
type myName name

type gram int
type kilogram gram
type ton kilogram

func main() {
	var you name
	var me myName

	fmt.Printf("you type %T\n", you)
	fmt.Printf("me type %T\n", me)

	//  ----  ==  -----

	var (
		salt   gram     = 100
		apples kilogram = 24
		truck  ton      = 10
	)

	fmt.Printf("\nsalt : %d, apples : %d, truck : %d\n", salt, apples, truck)

	// you need to converts all variables into the same type
	var combined = salt + gram(apples) + gram(truck)

	fmt.Printf("Combined variables : %d\n", combined)

	//  ----  ==  -----

	var (
		byteVal  byte
		uint8val uint8
	)

	// wont throw error because uint8val and byteVal has same underlying type uint8
	// note that byte is actually an uint8 type
	uint8val = byteVal

	// when it is already has the same underlying type
	// no need to converts to each other type
	// it will be unefficient
	uint8val = uint8(uint8val)
}
