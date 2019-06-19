package main

import (
	"fmt"

	dawlib "github.com/lego/first/version"
)

func main() {
	fmt.Print("Your Go Version is ")
	fmt.Println(dawlib.Version())
	dawlib.Thunder(5)
}
