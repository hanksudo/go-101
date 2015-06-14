package main

import (
	"fmt"
)

type Person struct {
	Name string
}

func (p *Person) Introduce() {
	fmt.Printf("Hi, I'm %s\n", p.Name)
}

type Saiyan struct {
	*Person
	Power int
}

func main() {
	goku := &Saiyan{
		Person: &Person{"Goku"},
		Power:  9001,
	}
	goku.Introduce()
}
