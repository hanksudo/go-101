// http://golang.org/doc/effective_go.html#interface_conversions
package main

import (
	"fmt"
)

type Stringer interface {
	String() string
}

func main() {
	var value interface{}
	value = "hello"

	fmt.Println(value.(string))

	switch str := value.(type) {
	case string:
		fmt.Println("it's string.", str)
	case Stringer:
		fmt.Println("it's Stringer.", str.String())
	}

	str, ok := value.(string)
	if ok {
		fmt.Printf("string value is: %q\n", str)
	} else {
		fmt.Printf("value is not a string\n")
	}
}
