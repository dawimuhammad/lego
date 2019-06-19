package main

import "fmt"

// Pyramid will print pyramid with parameter leng as pyramid height
func Pyramid(length int) {

	for i := 0; i < length; i++ {
		var rowString = ""

		for j := 0; j < (length - i); j++ {
			rowString += " "
		}

		for k := 0; k < ((2 * i) + 1); k++ {
			if i == (length-1) || i == 0 {
				rowString += "*"
			} else if (k == 0) || (k == (2 * i)) {
				rowString += "*"
			} else {
				rowString += " "
			}
		}

		fmt.Println(rowString)
	}
}
