package main

import (
	f "fmt"
	"runtime"
)

func main() {
	f.Println("Statements Package!")
	f.Println(runtime.NumCPU())
	f.Println(runtime.NumCPU() + 1)

	if 1 > 5 {
		f.Println("5 is higher than 1")
	} else {
		f.Println("1 is not higher than 5")
	}
}
