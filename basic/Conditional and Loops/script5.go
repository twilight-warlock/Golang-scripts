// 3 ways to define for loops
package main

import "fmt"
 
func main() {
 
	// 1
	k := 1
	for ; k <= 10; k++ {
		fmt.Println(k)
	}
 
	// 2
	k = 1
	for k <= 10 {
		fmt.Println(k)
		k++
	}
 
	// 3
	for k := 1; ; k++ {
		fmt.Println(k)
		if k == 10 {
			break
		}
	}
}