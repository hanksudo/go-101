// http://tour.golang.org/methods/8
package main

import (
	"fmt"
	"strconv"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

// built-in interface
// type error interface {
//     Error() string
// }
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}
func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}

	i, err := strconv.Atoi("a")
	if err != nil {
		fmt.Println("Couldn't convert number: ", err)
	}
	fmt.Println("Converted integer:", i)
}
