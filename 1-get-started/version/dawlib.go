// Learn Go Programming Course
// Creating your exported-package
//
// Copyright © 2018 dawimuhammad

package dawlib

import (
	"fmt"
	"runtime"
)

// Version function returns Go version number on your workspace
func Version() string {
	return runtime.Version()
}

// Thunder will return the biggest thunder from Thor son of Odin in Marvel Cinematic Universe
func Thunder(lightningLength int) {

	// for i := 0; i <= (lightningLength * 2); i++ {
	// 	fmt.Println("*")
	// }

	fmt.Println(`Thor thunder will coming in
5 ..
4 ..
3 ..
2 ..
1
	
THORRRRR!!!!`)

	for i := 0; i <= (lightningLength * 2); i++ {
		var stringLightning string

		if i == lightningLength {
			for j := 0; j < lightningLength; j++ {
				stringLightning += " *"
			}
		} else {
			var start int

			if i < lightningLength {
				start = i
			} else if i >= lightningLength {
				start = i - lightningLength
			}

			for k := 0; k < start; k++ {
				stringLightning += "  "
			}

			if i > 0 {
				stringLightning += "*"
			}
		}

		fmt.Println(stringLightning)
	}
}
