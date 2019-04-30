// Learn Go Programming Course
// Creating your exported-package
//
// Copyright © 2018 dawimuhammad

package dawlib

import (
	"runtime"
)

// Version function returns Go version number on your workspace
func Version() string {
	return runtime.Version()
}
