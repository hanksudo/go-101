// http://golang.org/doc/effective_go.html#interfaces_and_types
package main

import (
	"fmt"
	"sort"
)

type Sequence []int

func (s Sequence) Len() int {
	return len(s)
}
func (s Sequence) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s Sequence) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Method for printing - sorts the elements before printing
func (s Sequence) String() string {
	sort.Sort(s)
	str := "["
	for i, elem := range s {
		if i > 0 {
			str += " "
		}
		str += fmt.Sprint(elem)
	}
	return str + "]"
}
func main() {
	numbers := Sequence([]int{3, 1, 2})
	fmt.Println(numbers)
}
