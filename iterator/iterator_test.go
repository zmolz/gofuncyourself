package iterator_test

import (
	"testing"

	"github.com/zmolz/gofuncyourself/iterator"
)

func TestPipeline(t *testing.T) {
	t.Run("chained_pipeline_exhaust", func(t *testing.T) {
		r, _ := iterator.NewNumRange(1, 11, 1) // 1..10

		// Step 1 → Step 2: Filter even numbers
		even := iterator.NewFilterIterator(r, func(x int) bool { return x%2 == 0 })
		// even.Exhaust() = [2,4,6,8,10]

		// Step 2 → Step 3: Square numbers
		squared := iterator.NewMapIterator(even, func(x int) int { return x * x })
		// squared.Exhaust() = [4,16,36,64,100]

		// Step 3 → Step 4: Keep numbers > 20
		gt20 := iterator.NewFilterIterator(squared, func(x int) bool { return x > 20 })
		// gt20.Exhaust() = [36,64,100]

		// Step 4 → Step 5: Add 1
		final := iterator.NewMapIterator(gt20, func(x int) int { return x + 1 })
		// final.Exhaust() = [37,65,101]

		got := iterator.Exhaust(final)
		want := []int{37, 65, 101}

		if len(got) != len(want) {
			t.Fatalf("Exhaust() = %v, want %v", got, want)
		}
		for i := range got {
			if got[i] != want[i] {
				t.Errorf("Exhaust()[%d] = %v, want %v", i, got[i], want[i])
			}
		}
	})

	t.Run("chained_pipeline_next", func(t *testing.T) {
		r, _ := iterator.NewNumRange(1, 11, 1) // 1..10

		// Step 1 → Step 2: Filter even numbers
		even := iterator.NewFilterIterator(r, func(x int) bool { return x%2 == 0 })
		// even.Exhaust() = [2,4,6,8,10]

		// Step 2 → Step 3: Square numbers
		squared := iterator.NewMapIterator(even, func(x int) int { return x * x })
		// squared.Exhaust() = [4,16,36,64,100]

		// Step 3 → Step 4: Keep numbers > 20
		gt20 := iterator.NewFilterIterator(squared, func(x int) bool { return x > 20 })
		// gt20.Exhaust() = [36,64,100]

		// Step 4 → Step 5: Add 1
		final := iterator.NewMapIterator(gt20, func(x int) int { return x + 1 })
		// final.Exhaust() = [37,65,101]

		want := []int{37, 65, 101}
		var got int
		var ok bool
		for i := range 3 {
			got, _ = final.Next()
			if got != want[i] {
				t.Errorf("Next() = %v, want %v", got, want[i])
			}
		}

		if _, ok = final.Next(); ok {
			t.Fatalf("Next() return too many values")
		}
	})
}
