package main

// https://app.codility.com/programmers/lessons/2-arrays/odd_occurrences_in_array/

// []int{9, 3, 9, 3, 9, 7, 9}
// use mapping to rember match
// if exist
func Solution(A []int) int {
	if len(A) == 1 {
		return A[0]
	}

	pair := map[int]bool{}
	for _, n := range A {
		pair[n] = !pair[n]
	}
	for n, alone := range pair {
		if alone {
			return n
		}
	}
	return -1
}
