// Scope of variables
package main

import (
	"fmt"
)

// global variable

var s = "Japan"

func main() {
	fmt.Println(s)

	// local variable to main function
	x := true

	if x {
		// local variable to if braces
		y := 1
		if x != false {
			fmt.Println(s)
			fmt.Println(x)
			fmt.Println(y)
		}
	}
	fmt.Println(x)
	// fmt.Println(y) - undefined
}
