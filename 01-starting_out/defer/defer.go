package main 

import "fmt"

/*
		DEFER IN GO
	- The defer statement defers the execution of a function until the surrounding function returns. 
	
	- The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns. Deferred functions are executed in last-in-first-out order. This means that if you have multiple defer statements, they will be executed in reverse order of their appearance in the code.

	Defer is often used to ensure that resources are released, such as closing a file or unlocking a mutex, even if an error occurs. It can also be used to perform cleanup tasks or to log information when a function exits.
*/

func main() {
	fmt.Println("Counting...")

	for i:= 0; i < 10; i++ {
		defer fmt.Println(i) // defer the printing of i until the surrounding function (main) returns
	}
	
	fmt.Println("Done!")
}