// For range statement
package main

import "fmt"
 
func main() {
 
	// Example 1
	strDict := map[string]string{"Japan": "Tokyo", "China": "Beijing", "Canada": "Ottawa"}
	for index, element := range strDict {
		fmt.Println("Index :", index, " Element :", element)
	}
 
	// Printing keys
	for key := range strDict {
		fmt.Println(key)
	}
 
	// Printing value in keys
	for _, value := range strDict {
		fmt.Println(value)
	}

	// Looping over string i.e. length of string will become the times it loops, here 5 times
	for range "Hello" {
		fmt.Println("Hello")
	}
}