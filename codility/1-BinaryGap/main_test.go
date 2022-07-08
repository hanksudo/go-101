package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("9", func(t *testing.T) {
		got := Solution(9)
		want := 2

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, 9)
		}
	})

	t.Run("529", func(t *testing.T) {
		got := Solution(529)
		want := 4

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, 529)
		}
	})

	t.Run("20", func(t *testing.T) {
		got := Solution(20)
		want := 1

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, 20)
		}
	})
}
