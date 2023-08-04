package main

import (
	"fmt"
	"time"
)

func f(from string) {

	for i := 0; i < 3; i++ {
		time.Sleep(time.Duration(i*100) * time.Millisecond)
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	go f("goroutine")

	// start from an anonymous function
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
	fmt.Println(input)
}
