// Defered function calls, happens at the end
package main

import "fmt"
func first() {
	fmt.Println("First")
}
func second() {
	fmt.Println("Second")
}
func main() {
	defer second()
	first()
}

// Prints first then second