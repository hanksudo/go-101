// https://golang.org/doc/effective_go.html#panic
package main

import (
	"os"
)

func init() {
	var test = os.Getenv("TEST")
	if test == "" {
		panic("no value for $TEST")
	}
}
func main() {

}
