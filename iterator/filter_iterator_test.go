package iterator_test

import (
	"testing"

	"github.com/zmolz/gofuncyourself/iterator"
)

func TestFilter(t *testing.T) {
	t.Run("even_numbers_1_10", func(t *testing.T) {
		r, _ := iterator.NewNumRange(1, 11, 1) // 1..10

		isEven := func(i int) bool { return i%2 == 0 }

		filterIt := iterator.NewFilterIterator(r, isEven)

		got := iterator.Exhaust(filterIt)
		want := []int{2, 4, 6, 8, 10}

		if len(got) != len(want) {
			t.Fatalf("Exhaust() length = %v, want %v", len(got), len(want))
		}
		for i := range got {
			if got[i] != want[i] {
				t.Errorf("Exhaust()[%d] = %v, want %v", i, got[i], want[i])
			}
		}
	})

	t.Run("greater_than_5", func(t *testing.T) {
		r, _ := iterator.NewNumRange(1, 11, 1) // 1..10

		gt5 := func(i int) bool { return i > 5 }

		filterIt := iterator.NewFilterIterator(r, gt5)

		got := iterator.Exhaust(filterIt)
		want := []int{6, 7, 8, 9, 10}

		if len(got) != len(want) {
			t.Fatalf("Exhaust() length = %v, want %v", len(got), len(want))
		}
		for i := range got {
			if got[i] != want[i] {
				t.Errorf("Exhaust()[%d] = %v, want %v", i, got[i], want[i])
			}
		}
	})

	t.Run("no_matches", func(t *testing.T) {
		r, _ := iterator.NewNumRange(1, 5, 1) // 1..4

		never := func(i int) bool { return false }

		filterIt := iterator.NewFilterIterator(r, never)

		got := iterator.Exhaust(filterIt)
		want := []int{}

		if len(got) != len(want) {
			// t.Fatalf("Exhaust() length = %v, want %v", len(got), len(want))
			t.Fatalf("Exhaust() = %v, want %v", got, want)
		}
	})
}
