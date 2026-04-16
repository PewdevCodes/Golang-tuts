package main

import (
	"fmt"
	"time"
)

/* 
				GoRoutines

- Goroutines are lightweight threads managed by the Go runtime.
- They allow you to run functions concurrently, making it easier to perform tasks like API calls, data processing, and more without blocking the main thread.
- You can create a goroutine by using the `go` keyword followed by a function call.

*/

func calculateSum(numbers []int) {
    sum := 0
    for _, number := range numbers {
        sum += number
    }
    fmt.Println("Sum:", sum)
}

func main() {
	// Example of using goroutines for concurrent API calls
	go func() {
		fmt.Println("API call 1 started")
		time.Sleep(2 * time.Second)
		fmt.Println("API call 1 completed")
	}()

	go func() {
		fmt.Println("API call 2 started")
		time.Sleep(2 * time.Second)
		fmt.Println("API call 2 completed")
	}()

	time.Sleep(3 * time.Second)


	numbers := []int{1, 2, 3, 4, 5}
	go calculateSum(numbers)
	fmt.Println("Calculating sum in the background...")
	
}