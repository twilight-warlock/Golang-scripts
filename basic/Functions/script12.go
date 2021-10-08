// Recover function
package main

import "fmt"
func main() {
	var action int
    fmt.Println("Enter 1 for Student and 2 for Professional")
    fmt.Scanln(&action)
  
    switch action {
        case 1:
            fmt.Printf("I am a  Student")
        case 2:
            fmt.Printf("I am a  Professional")
		default:
			panic(fmt.Sprintf("I am a  %d",action))
    }    
    defer func() {
    	action := recover()    	
		fmt.Println(action)			
    }()
}

// Recover is a built-in function that regains control of a panicking goroutine. 
// Recover is only useful inside deferred functions. During normal execution, a call to recover will return nil and have no other effect. 
// If the current goroutine is panicking, a call to recover will capture the value given to panic and resume normal execution.