package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		got := Solution([]int{1, 3})
		want := 2

		if got != want {
			t.Errorf("got %d want %d given", got, want)
		}
	})

	t.Run("case 2", func(t *testing.T) {
		got := Solution([]int{1, 3, 6, 4, 2})
		want := 5

		if got != want {
			t.Errorf("got %d want %d given", got, want)
		}
	})

	t.Run("case 3", func(t *testing.T) {
		got := Solution([]int{})
		want := 1

		if got != want {
			t.Errorf("got %d want %d given", got, want)
		}
	})

	t.Run("case 4", func(t *testing.T) {
		got := Solution([]int{1, 2, 3})
		want := 4

		if got != want {
			t.Errorf("got %d want %d given", got, want)
		}
	})

	t.Run("case 5", func(t *testing.T) {
		got := Solution([]int{0})
		want := 1

		if got != want {
			t.Errorf("got %d want %d given", got, want)
		}
	})
}
