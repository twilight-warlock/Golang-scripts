// Declaring constants
package main

import "fmt"

const (
	PRODUCT  = "Mobile"
	QUANTITY = 50
	PRICE    = 50.50
	STOCK  = true
)

func main() {
	fmt.Println(QUANTITY)
	fmt.Println(PRICE)
	fmt.Println(PRODUCT)
	fmt.Println(STOCK)

	const abc int = 456
	// abc = 654   - cannot assign to abc (declared const)
	fmt.Println(abc)
}