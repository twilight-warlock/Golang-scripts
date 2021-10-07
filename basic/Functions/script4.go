// Passing pointers to functions
package main

import "fmt"

func update(a *int, t *string) {
	// defrencing pointer address
	*a = *a + 5      
	*t = *t + " Doe" 
	return
}

func main() {
	var age = 20
	var text = "John"
	fmt.Println("Before:", text, age)

	update(&age, &text)

	fmt.Println("After :", text, age)
}
