package main

// https://app.codility.com/programmers/lessons/3-time_complexity/perm_missing_elem/

func Solution(A []int) int {
	if len(A) == 0 {
		return 1
	}
	occur := map[int]bool{}
	for i := 0; i < len(A); i++ {
		occur[A[i]] = true
	}

	for i := 1; i <= len(A); i++ {
		if _, has := occur[i]; !has {
			return i
		}
	}
	return len(A) + 1
}
