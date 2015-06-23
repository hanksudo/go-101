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
