// Multiple global variable declaration and zero values
package main

import "fmt"

var (
	product  = "Mobile"
	quantity = 50
	price    = 50.50
	inStock  = true
)

func main() {
	fmt.Println(quantity)
	fmt.Println(price)
	fmt.Println(product)
	fmt.Println(inStock)

	// Zero value assigned to declarationns
	var unassignedVar float64 
	fmt.Println(unassignedVar)
}

/*

0 		- int
0     	- float32/64
"" 	  	- string
false 	- bool

*/
