package sliceutil_test

import (
	"fmt"
	"testing"

	"github.com/zmolz/gofuncyourself/sliceutil"
)

func TestMap(t *testing.T) {
	type test[T any, U any] struct {
		name string
		xs   []T
		f    func(T) U
		want []U
	}

	double := func(x int) int { return x * 2 }
	toString := func(x int) string { return fmt.Sprintf("%d", x) }
	identity := func(x int) int { return x }

	testsIntToInt := []test[int, int]{
		{
			name: "double_numbers",
			xs:   []int{1, 2, 3, 4},
			f:    double,
			want: []int{2, 4, 6, 8},
		},
		{
			name: "empty_slice",
			xs:   []int{},
			f:    double,
			want: []int{},
		},
		{
			name: "identity",
			xs:   []int{5, 10, 15},
			f:    identity,
			want: []int{5, 10, 15},
		},
	}

	for _, test := range testsIntToInt {
		t.Run(test.name, func(t *testing.T) {
			got := sliceutil.Map(test.xs, test.f)
			if len(got) != len(test.want) {
				t.Fatalf("Map() length = %v, want %v", len(got), len(test.want))
			}
			for i := range got {
				if got[i] != test.want[i] {
					t.Errorf("Map()[%d] = %v, want %v", i, got[i], test.want[i])
				}
			}
		})
	}

	testsIntToString := []test[int, string]{
		{
			name: "int_to_string",
			xs:   []int{1, 2, 3},
			f:    toString,
			want: []string{"1", "2", "3"},
		},
	}

	for _, test := range testsIntToString {
		t.Run(test.name, func(t *testing.T) {
			got := sliceutil.Map(test.xs, test.f)
			if len(got) != len(test.want) {
				t.Fatalf("Map() length = %v, want %v", len(got), len(test.want))
			}
			for i := range got {
				if got[i] != test.want[i] {
					t.Errorf("Map()[%d] = %v, want %v", i, got[i], test.want[i])
				}
			}
		})
	}
}
