package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		got := Solution([]int{9, 3, 9, 3, 9, 7, 9})
		want := 7

		if got != want {
			t.Errorf("got %d want %d given", got, want)
		}
	})

	t.Run("case 2", func(t *testing.T) {
		got := Solution([]int{1})
		want := 1

		if got != want {
			t.Errorf("got %d want %d given", got, want)
		}
	})
}
