package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"Lemon", "Melon", "Apple"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, -3}
	sort.Ints(ints)
	fmt.Println("Ints:", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted:", s)
}
