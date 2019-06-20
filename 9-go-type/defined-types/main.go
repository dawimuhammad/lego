package main

import (
	"fmt"
	"time"
)

// Duration defined type
type Duration int64

// Duration is the defined type from underlying type int64
// it has the same size as int64 type 8 bytes

type gram float64
type ounce float64

func main() {
	var ms Duration

	fmt.Printf("\n\n%T\n", ms)

	newTime, _ := time.ParseDuration("2h3m")

	fmt.Println(newTime.Minutes())

	// now := time.Now()
	lastTime := time.Date(2010, 12, 10, 20, 00, 00, 00, time.UTC)
	diff := time.Since(lastTime)

	// ---- === ----

	fmt.Println("time since : ", ((diff.Hours() / 24) / 365))
}
