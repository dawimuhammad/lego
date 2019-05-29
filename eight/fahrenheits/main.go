package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	celciusValue := os.Args[1]

	fmt.Printf("--- Celcius to Fahrenheit Conversion ---\n\n")
	fmt.Printf("Celcius value : %s C ,type %T\n", celciusValue, celciusValue)

	celcius, _ := strconv.ParseFloat(celciusValue, 64)

	fmt.Printf("Celcius formatted : %g C ,type %T\n", celcius, celcius)

	fahrenheit := celcius*1.8 + 32

	fmt.Printf("%g C equals to %g F\n", celcius, fahrenheit)
}
