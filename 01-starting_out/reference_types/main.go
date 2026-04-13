package main

import "fmt"

func main() { 
	i , j := 42, 2701 // variable declaration and assignment using short variable declaration syntax
	p, q := &i, &j // p and q are pointers to i and j respectively
	fmt.Println(*p, *q) // print the values pointed to by p and q
}