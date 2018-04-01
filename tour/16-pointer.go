/*
	Go has pointers. A pointer holds the memory address of a value.
	https://tour.golang.org/moretypes/1
*/
package main

import (
	"fmt"
)

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value for i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
	fmt.Println(*p)
}
