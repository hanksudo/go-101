package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		got := Solution(10, 85, 30)
		want := 3

		if got != want {
			t.Errorf("got %d want %d given", got, want)
		}
	})

	t.Run("case 2", func(t *testing.T) {
		got := Solution(1, 3, 1)
		want := 2

		if got != want {
			t.Errorf("got %d want %d given", got, want)
		}
	})
}
