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

	ch := make(chan int)
	go sum(a[:len(a)/2], ch)
	go sum(a[len(a)/2:], ch)

	x, y := <-ch, <-ch // receive from c
	// close(ch)
	fmt.Println(x, y, x+y) // 17 -5 12
}
