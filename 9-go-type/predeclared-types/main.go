package main

import (
	"fmt"
	"math"
)

func main() {
	var try1 uint8
	try1 = 255

	try2 := 255

	try1 += 10

	fmt.Printf("%d %d\n", try1, try2)

	if int(try1) < try2 {
		fmt.Println("try 2 is higher")
	}

	var muint8 uint8 = math.MaxUint8

	fmt.Println(muint8 + 10) // 9

	// fmt.Printf("%T %T %T\n", try1, try2, unit8)

	// floats

	var fl32 float32 = math.MaxFloat32

	fmt.Println(fl32 + 1)
}
