package main 

import (
	"fmt"
	"math"
)

/*
				CONDITIONS IN GO
	Go has the usual set of control structures, but there is no need for parentheses around conditions, and the braces are always required, even for single-line statements.
*/

func sqrt( x float64) string { // function to compute the square root of a number and return a string
	if x < 0 { // if the input is negative
		return sqrt(-x) + "i" // compute the square root of the positive value and append "i" to indicate an imaginary number
	}
	return fmt.Sprint(math.Sqrt(x)) // compute the square root of x and convert it to a string using fmt.Sprint
}

func pow( x , n , lim float64) float64 { // function to compute x raised to the power of n, but return lim if the result is greater than lim
	if v := math.Pow(x, n); v < lim { // compute x raised to the power of n and assign it to v, then check if v is less than lim
		return v // if v is less than lim, return v
	}	
	return lim // if v is not less than lim, return lim
}

func whichDay( day int) string {
	switch day { 
	case 1:
		return "Monday"
	case 2:
		return "Tuesday"
	case 3:	
		return "Wednesday"
	case 4:
		return "Thursday"
	case 5:	
		return "Friday"
	case 6:
		return "Saturday"
	case 7:
		return "Sunday"
	default:
		return "Invalid day"
	}
} 



func main() {
	fmt.Println(sqrt(2), sqrt(-4)) 

	fmt.Println(pow(3, 2, 10), pow(3, 3, 20)) 

	fmt.Println(whichDay(3)) 
	fmt.Println(whichDay(8))
}