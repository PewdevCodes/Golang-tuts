package main

import "fmt"

/*     

Structs are typed collections of fields. They are useful for grouping data together to form records. 
       
*/

type Vertex struct { 
	X , Y int 
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}       // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(Vertex{1, 2})

	v := Vertex{ 10 , 20 }
	v.X = v.X * 10
	v.Y = v.Y / 10
	fmt.Println(v)

	p := &v;
	p.X = 1
	fmt.Println(v)

}