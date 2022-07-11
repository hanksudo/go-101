package main

// https://app.codility.com/programmers/lessons/3-time_complexity/frog_jmp/

func Solution(X, Y, D int) int {
	ans := (Y - X) / D
	if (Y-X)%D > 0 {
		ans += 1
	}
	return ans
}
