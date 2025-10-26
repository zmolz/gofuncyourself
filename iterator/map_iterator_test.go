package iterator_test

import (
	"testing"

	"github.com/zmolz/gofuncyourself/iterator"
)

func TestMap(t *testing.T) {
	t.Run("square_num_range_1_5_1", func(t *testing.T) {
		r, _ := iterator.NewNumRange(1, 6, 1)

		square := func(i int) int { return i * i }

		mapIt := iterator.NewMapIterator(r, square)

		got := iterator.Exhaust(mapIt)
		want := []int{1, 4, 9, 16, 25}
		if len(got) != len(want) {
			t.Fatalf("Exhaust() length = %v, want %v", len(got), len(want))
		}
		for i := range got {
			if got[i] != want[i] {
				t.Errorf("Exhaust()[%d] = %v, want %v", i, got[i], want[i])
			}
		}

	})

	t.Run("id_num_range_1_5_1", func(t *testing.T) {
		r, _ := iterator.NewNumRange(5, 0, -1)

		id := func(i int) int { return i }

		mapIt := iterator.NewMapIterator(r, id)

		got := iterator.Exhaust(mapIt)
		want := []int{5, 4, 3, 2, 1}
		if len(got) != len(want) {
			t.Fatalf("Exhaust() length = %v, want %v", len(got), len(want))
		}
		for i := range got {
			if got[i] != want[i] {
				t.Errorf("Exhaust()[%d] = %v, want %v", i, got[i], want[i])
			}
		}

	})
}
