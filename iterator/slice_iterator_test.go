package iterator_test

import (
	"testing"

	. "github.com/zmolz/gofuncyourself/iterator"
)

func TestSliceIterator(t *testing.T) {
	t.Run("test_slice_iter", func(t *testing.T) {
		src := []int{1, 2, 3, 4}
		it := NewSliceIterator(src)

		got := Exhaust(it)
		want := []int{1, 2, 3, 4}

		if len(got) != len(want) {
			t.Fatalf("Exhaust() length = %v, want %v", len(got), len(want))
		}
		for i := range got {
			if got[i] != want[i] {
				t.Errorf("Exhaust()[%d] = %v, want %v", i, got[i], want[i])
			}
		}
	})

	t.Run("test_slice_empty", func(t *testing.T) {
		src := []int{}
		it := NewSliceIterator(src)

		got := Exhaust(it)
		want := []int{}

		if len(got) != len(want) {
			t.Fatalf("Exhaust() length = %v, want %v", len(got), len(want))
		}
	})
}
