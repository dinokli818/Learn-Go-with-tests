package main

import (
	"reflect"
	"testing"
)

func TestSumAll(t *testing.T) {
	t.Run("sum of an any size slices", func(t *testing.T) {
		numbers1 := []int{1, 2, 3, 4}
		numbers2 := []int{1, 3, 4}
		numbers3 := []int{1, 2, 3}
		got := SumAll(numbers1, numbers2, numbers3)
		want := []int{10, 8, 6}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("want '%v' but got '%v' given '%v %v %v'", want, got, numbers1, numbers2, numbers3)
		}
	})
}
