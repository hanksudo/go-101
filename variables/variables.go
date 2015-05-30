// define variables
package main

import "fmt"

var i int
var c, python, java bool

// If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
var j = 2
var ruby, golang, node = true, false, "no!"

// same as ruby, golan, node := true, false, "no!"

func main() {
	fmt.Println(i, c, python, java)
	fmt.Println(j, ruby, golang, node)
}
