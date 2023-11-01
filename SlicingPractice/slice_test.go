package slicePractice

import (
	"reflect"
	"testing"
)

func TestSlice(t *testing.T) {

	checkSlices := func(t testing.TB, got [][]int, want [][]int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	}
	t.Run("split a slice into half", func(t *testing.T) {
		got := SliceSplitHalf([]int{1, 2, 3, 4})
		want := [][]int{{1, 2}, {3, 4}}

		checkSlices(t, got, want)
	})

	t.Run("split a slice into number of slices", func(t *testing.T) {
		got := SliceSplitN([]int{1, 2, 3, 4, 5, 6}, 3)
		want := [][]int{{1, 2}, {3, 4}, {5, 6}}

		checkSlices(t, got, want)
	})
}
