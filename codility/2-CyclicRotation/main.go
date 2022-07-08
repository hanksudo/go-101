package main

// https://app.codility.com/programmers/lessons/2-arrays/cyclic_rotation/

/**
3, 8, 9, 7, 6 -> 9, 7, 6, 3, 8
3 [0:2] + [2:lenght]
8
*/
func Solution(A []int, K int) []int {
	if K <= 0 || len(A) == 0 || len(A) == K {
		return A
	}

	pos := len(A) - K%len(A)
	if pos > 0 {
		A = append(A[pos:], A[0:pos]...)
	}

	return A
}
