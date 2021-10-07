// If else
package main

import (
	"fmt"
)
 
func main() {
	x := 100
 
	if (x == 50) {
		fmt.Println("Germany")
	} else if x == 100 {
		fmt.Println("Japan")
	} else {
		fmt.Println("Canada")
	}
}

// No round braces in if statement also works, similar to python