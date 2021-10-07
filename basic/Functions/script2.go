// Why golang why
package main

import "fmt"

func cuboid(l int, b int, h int) (volume int) {
	var perimeter int
	perimeter =  4 * (l + b + h)
	fmt.Println("Parameter: ", perimeter)

	volume = l * b * h
	return // Return statement without specify variable name
}

func main() {
	fmt.Println("Volume: ", cuboid(20, 30, 50))
}