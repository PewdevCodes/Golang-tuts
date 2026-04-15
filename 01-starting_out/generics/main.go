package main

import "fmt"

/*  
			  GENERICS 
Generics allow us to write flexible and reusable code by enabling us to define functions, types, and data structures that can work with any data type. In Go, we can use type parameters to create generic functions and types. This allows us to write code that can operate on different types without sacrificing type safety.

*/
func Index[T comparable](s []T, v T) int {
	for i, item := range s {
		if item == v {
			return i
		}
	}
	return -1
}

func main() {
	strings := []string{"apple", "banana", "cherry"}
	fmt.Println(Index(strings, "banana")) 
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(Index(numbers, 3)) 
}