package main

import (
	"fmt"
)

type Vertex struct {
	X, Y int
}

var (
	p = Vertex{1, 2}  // has type Vertex
	q = &Vertex{1, 2} // has type *Vertext
	r = Vertex{X: 1}  // Y:0 is implicit
	s = Vertex{}      // X: and Y:0
)

func main() {
	fmt.Println(p, q, r, s)
}
