/*  
DATA TYPES IN GOLANG :: 

1) BASIC DATA TYPES ::
- bool
- string
- int, int8, int16, int32, int64
- uint, uint8, uint16, uint32, uint64
- float32, float64
- complex64, complex128

2) AGGREGATE DATA TYPES ::
- array ( COLLECTION OF ELEMENTS OF THE SAME TYPE WITH A FIXED SIZE )
- struct ( COLLECTION OF FIELDS OF DIFFERENT TYPES )

3) REFERENCE DATA TYPES ::
- pointer ( HOLDS THE MEMORY ADDRESS OF A VALUE )
- slice ( DYNAMICALLY SIZED, FLEXIBLE VIEW INTO THE ELEMENTS OF AN ARRAY )
- map ( COLLECTION OF KEY-VALUE PAIRS )
- channel ( USED FOR COMMUNICATION BETWEEN GOROUTINES )
- interface ( DEFINES A SET OF METHODS THAT A TYPE MUST IMPLEMENT TO SATISFY THE INTERFACE )
- function ( FIRST-CLASS CITIZENS IN Go, CAN BE ASSIGNED TO VARIABLES, PASSED AS ARGUMENTS, AND RETURNED FROM OTHER FUNCTIONS )



*/

package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("ToBe = %v\n", ToBe) // print the value of ToBe using the %v verb, which is a default format specifier for any value.
	fmt.Printf("MaxInt = %v\n", MaxInt) 
	fmt.Printf("z = %v\n", z)


/*    FORMAT SPECIFIERS IN Go's fmt package ::

	-- %v the value in a default format
	
	-- %T a Go-syntax representation of the type of the value
	
	-- %% a literal percent sign; consumes no value

	-- %q a double-quoted string safely escaped with Go syntax

	-- %b the binary representation of an integer

	-- %c the character represented by the corresponding Unicode code point

	-- %d the decimal representation of an integer

	-- %o the octal representation of an integer

	-- %x the hexadecimal representation of an integer, with a-f as 0-5

	-- %X the hexadecimal representation of an integer, with A-F as 0-5

	-- %U the Unicode format: U+1234; same as "U+%04X"
*/





}