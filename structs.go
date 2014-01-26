package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	fmt.Println(v)
	v.X = 4
	fmt.Println(v)

	// pointer, will replace v.X
	q := &v
	q.X = 1e9
	fmt.Println(v)

	// non-pointer
	p := v
	p.X = 123
	fmt.Println(v)

	// new function - allocates zero value and return pointer
	u := new(Vertex) // var u *U = new(Vertex)
	fmt.Println(u)
	u.X, u.Y = 11, 9
	fmt.Println(u)
}
