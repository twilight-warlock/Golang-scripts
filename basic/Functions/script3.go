// Multiple return values
package main

import "fmt"

func rectangle(l int, b int) (area int, perimeter int) {
	perimeter = 2 * (l + b)
	area = l * b
	return // Return statement without specify variable name
}

func main() {
	var a, p int
	a, p = rectangle(20, 30)
	fmt.Println("Area:", a)
	fmt.Println("Perimeter:", p)
}