package main

import (
	"reflect"
	"testing"
)

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(want, got) {
			t.Errorf("want '%v' but got '%v' ", want, got)
		}
	}
	t.Run("sum of an any size slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3, 4}, []int{1, 3, 4}, []int{1, 2, 3})
		want := []int{9, 7, 5}
		checkSums(t, got, want)
	})
	t.Run("sum of an any size slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1}, []int{1, 2})
		want := []int{0, 0, 2}
		checkSums(t, got, want)
	})

}
