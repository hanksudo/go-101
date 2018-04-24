// http://tour.golang.org/concurrency/2
package main

import (
	"fmt"
)

func sum(a []int, ch chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	ch <- sum // send sum to c
}
func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)

	x, y := <-c, <-c // receive from c
	fmt.Println(x, y, x+y)
}
