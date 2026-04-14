package main

import "fmt"

/*

Slices are a key data type in Go, giving a more powerful interface to sequences than arrays. 
A slice is a descriptor of an array segment. It consists of a pointer to the array, the length of the segment, and its capacity (the maximum length of the segment).



Slices = struct { 
	array pointer
	length
	capacity
}

*/

func main() {

	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4] // [ start : end ]  end is not included
	fmt.Println(s)

	names := [4]string{
		"john",
		"paul",
		"george",
		"ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b = append(b, "x")
	fmt.Println(a, b)

	// declaration of a slice without an array

	var s1 []int
	fmt.Println(s1, len(s1), cap(s1))

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s3 := []struct {
	i int
	b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s3)


	a2 := make([]int, 5) // len(a)=5
	fmt.Println(a2)


	// Append on Slices

	var s4 []int
	printSlice(s4)

	s4 = append(s4, 0)
	printSlice(s4)

	s4 = append(s4, 1)
	printSlice(s4)

	s4 = append(s4, 2, 3, 4)
	printSlice(s4)

	var numbers[] int;
	for i := 0; i < 20; i++ {
		numbers = append(numbers, i)
	}
	printSlice(numbers)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

}
