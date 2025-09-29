package sliceutil_test

import (
	"errors"
	"testing"

	"github.com/zmolz/gofuncyourself/sliceutil"
)

func TestFoldLeft(t *testing.T) {
	type test struct {
		name string
		xs   []int
		a    int
		f    func(int, int) int
		want int
	}

	sumFunc := func(x, acc int) int { return x + acc }
	subtractFunc := func(x, acc int) int { return acc - x }
	productFunc := func(x, acc int) int { return x * acc }

	tests := []test{
		{
			name: "sum_foldleft",
			xs:   []int{1, 2, 3, 4},
			a:    0,
			f:    sumFunc,
			want: 10,
		},
		{
			name: "product_foldleft",
			xs:   []int{1, 2, 3, 4},
			a:    1,
			f:    productFunc,
			want: 24,
		},
		{
			name: "subtract_foldleft_non_commutative",
			xs:   []int{10, 1, 2},
			a:    0,
			f:    subtractFunc,
			want: -13, // (((0-10)-1)-2) = -13
		},
		{
			name: "empty_slice_foldleft",
			xs:   []int{},
			a:    42,
			f:    sumFunc,
			want: 42,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := sliceutil.FoldLeft(test.xs, test.a, test.f)
			if got != test.want {
				t.Errorf("FoldLeft() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestReduce(t *testing.T) {
	type test struct {
		name    string
		xs      []int
		f       func(int, int) int
		want    int
		wantErr error
	}

	maxFunc := func(x, acc int) int {
		if acc > x {
			return acc
		}
		return x
	}
	sumFunc := func(x, acc int) int {
		return x + acc
	}
	productFunc := func(x, acc int) int {
		return x * acc
	}
	subtractFunc := func(x, acc int) int {
		return acc - x
	}

	tests := []test{
		{
			name:    "max_filled_slice",
			xs:      []int{1, 2, 3, 4},
			f:       maxFunc,
			want:    4,
			wantErr: nil,
		},
		{
			name:    "max_empty_slice",
			xs:      []int{},
			f:       maxFunc,
			want:    0,
			wantErr: sliceutil.ErrEmptyReduce,
		},
		{
			name:    "max_one_element",
			xs:      []int{10},
			f:       maxFunc,
			want:    10,
			wantErr: nil,
		},
		{
			name:    "sum_of_integers",
			xs:      []int{1, 2, 3, 4},
			f:       sumFunc,
			want:    10,
			wantErr: nil,
		},
		{
			name:    "product_of_integers",
			xs:      []int{1, 2, 3, 4},
			f:       productFunc,
			want:    24,
			wantErr: nil,
		},
		{
			name:    "sum_with_negatives",
			xs:      []int{-1, -2, -3, -4},
			f:       sumFunc,
			want:    -10,
			wantErr: nil,
		},
		{
			name:    "non_commutative_subtraction",
			xs:      []int{10, 1, 2},
			f:       subtractFunc,
			want:    7, // (10 - 1) - 2 = 7
			wantErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := sliceutil.Reduce(test.xs, test.f)

			if !errors.Is(err, test.wantErr) {
				t.Errorf("Reduce() errored with %v, expected %v", err, test.wantErr)
			}
			if got != test.want {
				t.Errorf("Reduce() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestFoldRight(t *testing.T) {
	type test struct {
		name string
		xs   []int
		a    int
		f    func(int, int) int
		want int
	}

	sumFunc := func(x, acc int) int { return x + acc }
	subtractFunc := func(x, acc int) int { return x - acc }
	productFunc := func(x, acc int) int { return x * acc }

	tests := []test{
		{
			name: "sum_foldright",
			xs:   []int{1, 2, 3, 4},
			a:    0,
			f:    sumFunc,
			want: 10,
		},
		{
			name: "product_foldright",
			xs:   []int{1, 2, 3, 4},
			a:    1,
			f:    productFunc,
			want: 24,
		},
		{
			name: "subtract_foldright_non_commutative",
			xs:   []int{10, 1, 2},
			a:    0,
			f:    subtractFunc,
			want: 11, // (10 - (1 - (2 - 0))) = 11
		},
		{
			name: "empty_slice_foldright",
			xs:   []int{},
			a:    42,
			f:    sumFunc,
			want: 42,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := sliceutil.FoldRight(test.xs, test.a, test.f, test.name == "subtract_foldright_non_commutative")
			if got != test.want {
				t.Errorf("FoldRight() = %v, want %v", got, test.want)
			}
		})
	}
}
