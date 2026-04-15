package main

import (
	"fmt"
	"math"
)

func complete(fn func(float64 , float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypnot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Printf("The hypotenuse of a right triangle with sides 3 and 4 is: %.2f\n", complete(hypnot))
}

