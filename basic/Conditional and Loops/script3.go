// Switch case - all types
package main

import (
	"fmt"
	"time"
)
 
func main() {
	today := time.Now()
	var t int = today.Day()
 
	switch t {
	case 1:
		fmt.Println("All th best for the month ahead!")
	case 5, 10, 15:
		fmt.Println("Clean your house.")
	case 25, 26, 27:
		fmt.Println("Buy some food.")
		fallthrough
	case 31:
		fmt.Println("Party tonight.")
	default:
		fmt.Println("No information available for that day.")
	}
}