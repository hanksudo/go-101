package main

// https://app.codility.com/programmers/lessons/3-time_complexity/tape_equilibrium/

import (
	"math"
)

func Solution(A []int) int {
	ans := math.MaxUint32
	prev := 0
	total := 0
	for _, v := range A {
		total += v
	}
	for i := 0; i < len(A)-1; i++ {
		prev += A[i]
		next := total - prev
		diff := next - prev
		if diff < 0 {
			diff *= -1
		}
		if diff < ans {
			ans = diff
		}
	}
	return ans
}
