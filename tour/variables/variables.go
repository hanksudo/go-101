// define variables
package main

import "fmt"

var i int
var c, python, java bool

var april, janurary int = 4, 1

// If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
var j = 2
var ruby, golang, node = true, false, "no!"

func main() {
	// define without var and type, this can only use inside of functions
	angular, backbone, ember := true, false, "no!"
	_, may := 1, 5

	fmt.Println(i, c, python, java)
	fmt.Println(april, janurary)
	fmt.Println(j, ruby, golang, node)
	fmt.Println(angular, backbone, ember)
	fmt.Println(may)
}
