package main

import "fmt"

// Vertex struct
type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

/*
To access the field X of a struct when we have the struct pointer p we could write (*p).X.
However, that notation is cumbersome, so the language permits us instead to write just p.X,
without the explicit dereference.
*/
func main() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println("X coordinates are : ", v.X)
	fmt.Println("Vertex is: ", v1, "\nPointer to vetex: ", p, "\nvertex is with Y =0 : ",
		v2, "\nvertex is with X and Y both 0: ", v3)
}
