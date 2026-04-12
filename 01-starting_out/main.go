package main // packages for organising code and controlling visibility
// -- package names must be unique and lowercase, and should not contain underscores or hyphens and should be short and descriptive. The main package is a special package that serves as the entry point for the program. It must contain a main function, which is the starting point of the program execution.

import (
	"fmt" // import the fmt package for formatted I/O
	"math" // import the math package for mathematical 		  functions
	"math/rand" // import the rand package for generating random numbers

 ) 
// func add( a , b int) int (return type) { 
// 	return a + b
// }
 func add(a int , b int) int { // function to add two integers and return the result
	return a + b
}

func swap( x , y string) ( string , string) {
	return y , x
}

func split( sum int) ( x int, y int) { 
	x = sum * 4 / 9
	y = sum - x
	return // return the values of x and y (naked return)
}

func main() { // entry point of the program
	fmt.Println("Hello, World!")
	fmt.Println("My fav number is :" , rand.Intn(10)) // print a random number between 0 and 9
	fmt.Println("The square root of 7 is:", math.Sqrt(7)) // print the square root of 7
	fmt.Println("The value of Pi is:", math.Pi) // print the value of Pi

	fmt.Println(add(43,34)) // call the add function with arguments 43 and 34, and print the result

	fmt.Println(swap("Hello", "World")) // call the swap function with arguments "Hello" and "World", and print the result	

	s, r := swap("Joey", "Pewd") // call the swap function and assign the results to variables s and r ( variable declaration and assignment using short variable declaration syntax)
	fmt.Println(s, r) // print the values of s and r

	fmt.Println(split(79))
} 