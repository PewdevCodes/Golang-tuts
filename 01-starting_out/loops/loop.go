package main 

import "fmt"

/*
		LOOPS IN GO
	Go has only one looping construct, the for loop. The basic syntax of a for loop is as follows:

	for initialization; condition; post {
		// loop body
	}
*/

func main() {
	sum := 0 
	for i := 0 ; i < 10 ; i++ { // for loop with initialization, condition, and post statement
		sum += i // add the value of i to sum
	}
	fmt.Println(sum) // print the final value of sum

	// making while loop using for loop
	newsum := 1
	for newsum < 1000 { // for loop with only a condition, acting as a while loop
		newsum += newsum // double the value of newsum
	}
	fmt.Println(newsum) // print the final value of newsum
}