package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("sum of an any size slices", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4}
		got := Sum(numbers)
		want := 10
		if got != want {
			t.Errorf("want '%d' but got '%d' given '%v'", want, got, numbers)
		}
	})
}
