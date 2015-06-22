package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	ch := make(chan int, 10)
	go fibonacci(cap(ch), ch)

	// test whether chanel has been closed
	v, ok := <-ch
	fmt.Println(v, ok)

	// receives values from channel until it's closed
	for i := range ch {
		fmt.Println(i)
	}
}
