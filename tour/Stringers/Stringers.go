// http://tour.golang.org/methods/6
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}
func main() {
	a := Person{"Hank Wang", 30}
	z := Person{"Zara", 9001}
	fmt.Println(a, z)
}
