package main

import (
	"reflect"
	"testing"
)

func TestQuicksort(t *testing.T) {
	t.Run("Test case 1", func(t *testing.T) {
		t.Helper()
		got := []int{3, 7, 2, 9, 4}
		Quicksort(got, 0, len(got)-1)
		want := []int{2, 3, 4, 7, 9}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got error %q want %q", got, want)
		}
	})

	t.Run("Test case 2", func(t *testing.T) {
		t.Helper()
		got := []int{2, 1, 4, 13, 14, 12, 3, 16, 5, 2, 10}
		Quicksort(got, 0, len(got)-1)
		want := []int{1, 2, 2, 3, 4, 5, 10, 12, 13, 14, 16}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got error %q want %q", got, want)
		}
	})
}
