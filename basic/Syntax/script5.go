// Short Variable Declaration
package main

import (
	"fmt"
	"reflect"
)

func main() {
	name := false
	fmt.Println(reflect.TypeOf(name))
}

// Easier and shortcut way to declare variables with initialization