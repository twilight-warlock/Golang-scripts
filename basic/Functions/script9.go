// Variadic functions
package main

import "fmt"

func main() {
	
	variadicExample()
	variadicExample("red", "blue")
	variadicExample("red", "blue", "green")
	variadicExample("red", "blue", "green", "yellow")
}

func variadicExample(s ...string) {
	fmt.Println(s)
	if len(s)>0 {
		fmt.Println("Last element:",s[len(s)-1])
	}
}