package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "UPPER CASE STRINGS"
	lower := strings.ToLower(str)
	fmt.Println("lower string:", lower)

	hello := "Hello, 世界！"
	hello2 := [5]byte{'h', 'e', 'l', 'l', 'o'}

	fmt.Printf("%s\n", hello)
	fmt.Printf("%s\n", hello2)

	// combine two strings
	a := "Hello, "
	b := "world!"
	c := a + b
	fmt.Println(c)

	// replace one character
	byteHello := []byte(hello) // convert string to []byte type
	byteHello[0] = 'h'
	newHello := string(byteHello)
	fmt.Printf("%s\n", newHello)

	fmt.Printf("%s\n", newHello[7:])

	// multi line string
	multi := `Hello
    世界！`
	fmt.Println(multi)

	fmt.Println(string(77)) // M
}
