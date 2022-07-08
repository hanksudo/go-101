package main

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		got := Solution([]int{3, 8, 9, 7, 6}, 3)
		want := []int{9, 7, 6, 3, 8}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d given", got, want)
		}
	})

	t.Run("case 2", func(t *testing.T) {
		got := Solution([]int{0, 0, 0}, 1)
		want := []int{0, 0, 0}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d given", got, want)
		}
	})

	t.Run("case 3", func(t *testing.T) {
		got := Solution([]int{1, 2, 3, 4}, 4)
		want := []int{1, 2, 3, 4}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d given", got, want)
		}
	})
}
