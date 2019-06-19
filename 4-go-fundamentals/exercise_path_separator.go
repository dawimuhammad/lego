package main

import (
	"fmt"
	"path"
)

// PathSeparator will return directory path and filename from given filepath string
func PathSeparator(filepath string) {
	dir, filename := path.Split(filepath)

	fmt.Println("directory : ", dir,
		"\nfilename : ", filename)
}
