package main

// https://app.codility.com/programmers/lessons/1-iterations/binary_gap/

func Solution(N int) int {
	ans := 0
	last := -1
	pos := 0
	for {
		if N == 0 {
			return ans
		}
		m := N % 2
		if m == 1 {
			if last == -1 {
				last = pos
			} else {
				diff := pos - last - 1
				if diff > ans {
					ans = diff
				}
				last = pos
			}
		}
		N /= 2
		pos += 1
	}
}
