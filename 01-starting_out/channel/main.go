package main

import (
	"fmt"
)
/*  
		Think of it like this:

		Two people passing a ball
		One throws (ch <-)
		One catches (<-ch)


*/
//  c: = make(chan int) 
// Create a channel to communicate between goroutines in integer type
// You can send and receive data from a channel using the <- operator.

func main() {
  ch := make(chan string)

 go func() {
  ch <- "Hello from goroutine"
 }()

 message := <-ch
 fmt.Println(message)
}