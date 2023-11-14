package main

import (
	"reflect"
	"testing"
)

func TestArraySum(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := ArraySum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	var numbers1 = []int{1, 2, 3}
	var numbers2 = []int{4, 5}
	got := SumAll(numbers1, numbers2)
	want := []int{6, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d want %d given %v %v", got, want, numbers1, numbers2)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSum := func(t testing.TB, got []int, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("make sum of slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{2, 3, 6})
		want := []int{2, 9}

		checkSum(t, got, want)
	})

	t.Run("safely handle empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{2, 3, 6})
		want := []int{0, 9}

		checkSum(t, got, want)
	})

}
