package sliceutil_test

import (
	"testing"

	"github.com/zmolz/gofuncyourself/sliceutil"
)

func TestFlatten(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		xss := [][]int{}
		want := []int{}
		got := sliceutil.Flatten(xss)

		if len(got) != len(want) {
			t.Fatalf("Flatten() length = %v, want %v", len(got), len(want))
		}
		for i := range got {
			if got[i] != want[i] {
				t.Errorf("Flatten()[%d] = %v, want %v", i, got[i], want[i])
			}
		}
	})

	t.Run("flatten_singletons", func(t *testing.T) {
		xss := [][]int{{1}, {2}, {3}}
		want := []int{1, 2, 3}
		got := sliceutil.Flatten(xss)

		if len(got) != len(want) {
			t.Fatalf("Flatten() length = %v, want %v", len(got), len(want))
		}
		for i := range got {
			if got[i] != want[i] {
				t.Errorf("Flatten()[%d] = %v, want %v", i, got[i], want[i])
			}
		}
	})

	t.Run("flatten_triangle", func(t *testing.T) {
		xss := [][]int{{1}, {2, 3}, {4, 5, 6}}
		want := []int{1, 2, 3, 4, 5, 6}
		got := sliceutil.Flatten(xss)

		if len(got) != len(want) {
			t.Fatalf("Flatten() length = %v, want %v", len(got), len(want))
		}
		for i := range got {
			if got[i] != want[i] {
				t.Errorf("Flatten()[%d] = %v, want %v", i, got[i], want[i])
			}
		}
	})
}

// func TestArbitraryFlatten(t *testing.T) {
// 	t.Run("flatten_triple_triangle", func(t *testing.T) {
// 		xss := [][][]int{
// 			{
// 				{1},
// 			},
// 			{
// 				{2, 3},
// 				{4, 5, 6},
// 			},
// 			{
// 				{7, 8, 9, 10},
// 				{11, 12, 13, 14, 15},
// 				{16, 17, 18, 19, 20, 21},
// 			},
// 		}

// 		want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21}
// 		got := sliceutil.ArbitraryFlatten(xss)

// 		if len(got) != len(want) {
// 			t.Fatalf("Flatten() length = %v, want %v", len(got), len(want))
// 		}
// 		for i := range got {
// 			if got[i] != want[i] {
// 				t.Errorf("Flatten()[%d] = %v, want %v", i, got[i], want[i])
// 			}
// 		}
// 	})
// }
