// http://golang.org/doc/effective_go.html#conversions
package main

import (
	"fmt"
	"sort"
)

type Sequence []int

// Method for printing - sorts the elements before printing
func (s Sequence) String() string {
	sort.IntSlice(s).Sort()
	return fmt.Sprint([]int(s))
}
func main() {
	numbers := Sequence([]int{3, 1, 2})
	fmt.Println(numbers)
}
