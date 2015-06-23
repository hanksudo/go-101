package main

import (
	"fmt"
	"os"
)

type T2 struct {
	a int
	b float64
	c string
}

func (t *T2) String() string {
	return fmt.Sprintf("%d/%g/%q", t.a, t.b, t.c)
}

func main() {
	fmt.Printf("Hello %d\n", 23)
	fmt.Fprint(os.Stdout, "Hello ", 23, "\n")
	fmt.Println("Hello", 23)
	fmt.Println(fmt.Sprint("Hello ", 23))

	var x uint64 = 1<<64 - 1
	fmt.Printf("%d %x; %d %x\n", x, x, int64(x), int64(x))

	// print any value
	s := "yo"
	fmt.Printf("%v\n", s)
	fmt.Println(s)

	// print struct
	type T struct {
		a int
		b float64
		c string
	}
	t := &T{7, 3.14, "abc\tdef"}
	fmt.Printf("%v\n", t)
	fmt.Printf("%+v\n", t)
	fmt.Printf("%#v\n", t)

	// print type of value
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", t)

	// custom format
	t2 := &T2{7, 3.14, "abc\tdef"}
	fmt.Printf("%v\n", t2)
}
