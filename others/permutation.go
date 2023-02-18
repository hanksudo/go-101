// Print all permutations of a slice characters or string
package main

import (
	"fmt"
)

// Implement the perm() function that accepts a slice or string and prints all possible combinations of characters.
func Perm(a []rune, f func([]rune)) {
	perm(a, f, 0)
}
func perm(a []rune, f func([]rune), i int) {
	if i > len(a) {
		f(a)
		return
	}

	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

func main() {
	Perm([]rune("abc"), func(a []rune) {
		fmt.Println(string(a))
	})
}

/*
abc
acb
bac
bca
cba
cab
*/
