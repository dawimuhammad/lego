package main

import "fmt"

func main() {
	fmt.Println(2 + 4 - 3*3/2)

	// same as

	fmt.Println(2 + 4 - ((3 * 3) / 2))

	// precedence operators order
	// same level perform from the left to the right
	// multiply & divide
	// add & substract
}
