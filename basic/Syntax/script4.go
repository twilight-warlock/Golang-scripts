// Declaration of multiple variables
package main

import (
	"fmt"
)

func main() {
	var fname, lname string = "John", "Doe" 
	m, n, o := 1, 2, 3 
	item, price := "Mobile", 2000

	fmt.Println(fname +" "+ lname)
	fmt.Println(m + n + o)
	fmt.Println(item, "-", price)
}

/* 
	1) +" "+ === ,
	2) Multiple declarations by using new line starting with variable name followed by it's value
	3) For first line variables, we use = but for the following lines we use :=
*/ 