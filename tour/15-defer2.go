package main

import (
	"fmt"
)

func start(s string) string {
	fmt.Println("starting:", s)
	return s
}

func end(s string) {
	fmt.Println("ending:", s)
}

func a() {
	defer end(start("a"))
	fmt.Println("in a")
	b()
}

func b() {
	defer end(start("b"))
	fmt.Println("in b")
}
func main() {
	a()
}
