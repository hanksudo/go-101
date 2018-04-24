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
	// Closing is only necessary when the receiver must be told
	// there are no more values coming, such as to terminate ta range loop.
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)

	// test whether chanel has been closed
	v, ok := <-c
	fmt.Println(v, ok)

	// receives values from channel until it's closed
	for i := range c {
		fmt.Println(i)
	}
}
