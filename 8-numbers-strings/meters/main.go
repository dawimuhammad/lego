package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Printf("\nConvert your feet value to meters\n")

	feetValue := os.Args[1]

	fmt.Printf("Input value is %s feet with type %T\n", feetValue, feetValue)

	feet, _ := strconv.ParseFloat(feetValue, 64)

	fmt.Printf("Formatted feet value is %g with type %T\n", feet, feet)

	meters := feet * 0.3048

	fmt.Printf("%g feet equals to %g meters\n", feet, meters)
}
