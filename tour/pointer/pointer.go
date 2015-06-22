package main

import (
	"fmt"
)

func main() {
	i, j := 42, 2701
	p := &i
	*p = 21
	fmt.Println(i)
	fmt.Println(*p)
	// fmt.Println(&p)

	p = &j
	*p = *p / 37
	fmt.Println(j)
	fmt.Println(*p)
}
