package iterator_test

import (
	"testing"

	"github.com/zmolz/gofuncyourself/iterator"
)

func TestNumRange(t *testing.T) {
	t.Run("test_1_11_2", func(t *testing.T) {
		rng, err := iterator.NewNumRange(1, 11, 2)

		if err != nil {
			t.Fatalf("Error Creating Iterator")
		}

		got := iterator.Exhaust(rng)
		want := []int{1, 3, 5, 7, 9}

		if len(got) != len(want) {
			t.Fatalf("Exhaust() length = %v, want %v", len(got), len(want))
		}
		for i := range got {
			if got[i] != want[i] {
				t.Errorf("Exhaust()[%d] = %v, want %v", i, got[i], want[i])
			}
		}
	})

	t.Run("test_wrong_dir", func(t *testing.T) {
		rng, err := iterator.NewNumRange(11, 0, 2)

		if err != nil {
			t.Fatalf("Error Creating Iterator")
		}

		got := iterator.Exhaust(rng)
		want := []int{}

		if len(got) != len(want) {
			t.Fatalf("Exhaust() length = %v, want %v", len(got), len(want))
		}
	})

	t.Run("test_neg_dir", func(t *testing.T) {
		rng, err := iterator.NewNumRange(11, 0, -2)

		if err != nil {
			t.Fatalf("Error Creating Iterator")
		}

		got := iterator.Exhaust(rng)
		want := []int{11, 9, 7, 5, 3, 1}
		
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
