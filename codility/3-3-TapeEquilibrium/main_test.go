package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		got := Solution([]int{3, 1, 2, 4, 3})
		want := 1

		if got != want {
			t.Errorf("got %d want %d given", got, want)
		}
	})
}
