package sliceutil_test

import (
	"testing"

	"github.com/zmolz/gofuncyourself/sliceutil"
)

func TestFilterMap(t *testing.T) {
	type test struct {
		name string
		xs   []int
		want []int
		f    func(int) (int, bool)
	}

	tests := []test{
		{
			name: "filtermap_square_odds",
			xs:   []int{1, 2, 3, 4, 5},
			want: []int{1, 9, 25},
			f: func(x int) (int, bool) {
				if x%2 == 0 {
					return 0, false
				}

				return x * x, true
			},
		},
		{
			name: "filtermap_none_true",
			xs:   []int{1, 2, 3, 4, 5},
			want: []int{},
			f: func(x int) (int, bool) {
				if x != 10 {
					return 0, false
				}

				return x * x, true
			},
		},
		{
			name: "filtermap_empty",
			xs:   []int{},
			want: []int{},
			f: func(x int) (int, bool) {
				return 0, false
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := sliceutil.FilterMap(test.xs, test.f)
			if len(got) != len(test.want) {
				t.Fatalf("Flatten() length = %v, want %v", len(got), len(test.want))
			}
			for i := range got {
				if got[i] != test.want[i] {
					t.Errorf("Flatten()[%d] = %v, want %v", i, got[i], test.want[i])
				}
			}
		})
	}
}

func TestFlatMap(t *testing.T) {
	type test struct {
		name string
		xs   []int
		want []int
		f    func(int) []int
	}

	tests := []test{
		{
			name: "flatmap_duplicate_each",
			xs:   []int{1, 2, 3},
			want: []int{1, 1, 2, 2, 3, 3},
			f: func(x int) []int {
				return []int{x, x}
			},
		},
		{
			name: "flatmap_expand_to_range",
			xs:   []int{1, 2, 3},
			want: []int{1, 1, 2, 1, 2, 3},
			f: func(x int) []int {
				out := []int{}
				for i := 1; i <= x; i++ {
					out = append(out, i)
				}
				return out
			},
		},
		{
			name: "flatmap_to_empty",
			xs:   []int{1, 2, 3},
			want: []int{},
			f: func(x int) []int {
				return []int{}
			},
		},
		{
			name: "flatmap_empty_input",
			xs:   []int{},
			want: []int{},
			f: func(x int) []int {
				return []int{x * 2}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := sliceutil.FlatMap(test.xs, test.f)
			if len(got) != len(test.want) {
				t.Fatalf("FlatMap() length = %v, want %v", len(got), len(test.want))
			}
			for i := range got {
				if got[i] != test.want[i] {
					t.Errorf("FlatMap()[%d] = %v, want %v", i, got[i], test.want[i])
				}
			}
		})
	}
}
