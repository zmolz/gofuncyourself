package sliceutil_test

import (
	"errors"
	"testing"

	"github.com/zmolz/gofuncyourself/sliceutil"
)

func TestRemove(t *testing.T) {
	type test[T any] struct {
		name    string
		xs      []T
		i       int
		want    []T
		wantErr error
	}

	tests := []test[int]{
		{
			name:    "remove_middle",
			xs:      []int{1, 2, 3, 4, 5},
			i:       2,
			want:    []int{1, 2, 4, 5},
			wantErr: nil,
		},
		{
			name:    "remove_first",
			xs:      []int{1, 2, 3},
			i:       0,
			want:    []int{2, 3},
			wantErr: nil,
		},
		{
			name:    "remove_last",
			xs:      []int{1, 2, 3},
			i:       2,
			want:    []int{1, 2},
			wantErr: nil,
		},
		{
			name:    "single_element_slice",
			xs:      []int{42},
			i:       0,
			want:    []int{},
			wantErr: nil,
		},
		{
			name:    "negative_index",
			xs:      []int{1, 2, 3},
			i:       -1,
			want:    nil,
			wantErr: sliceutil.ErrIndexOutOfRange,
		},
		{
			name:    "index_too_large",
			xs:      []int{1, 2, 3},
			i:       3,
			want:    nil,
			wantErr: sliceutil.ErrIndexOutOfRange,
		},
		{
			name:    "empty_slice",
			xs:      []int{},
			i:       0,
			want:    nil,
			wantErr: sliceutil.ErrIndexOutOfRange,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := sliceutil.Remove(test.xs, test.i)

			if !errors.Is(err, test.wantErr) {
				t.Fatalf("Remove() error = %v, want %v", err, test.wantErr)
			}

			if err == nil {
				if len(got) != len(test.want) {
					t.Fatalf("Remove() length = %v, want %v", len(got), len(test.want))
				}
				for j := range got {
					if got[j] != test.want[j] {
						t.Errorf("Remove()[%d] = %v, want %v", j, got[j], test.want[j])
					}
				}
			}
		})
	}
}

func TestRemoveUnordered(t *testing.T) {
	type test[T any] struct {
		name    string
		xs      []T
		i       int
		want    []T
		wantErr error
	}

	tests := []test[int]{
		{name: "remove_middle", xs: []int{1, 2, 3, 4, 5}, i: 2, want: []int{1, 2, 5, 4}, wantErr: nil},
		{name: "remove_first", xs: []int{1, 2, 3}, i: 0, want: []int{3, 2}, wantErr: nil},
		{name: "remove_last", xs: []int{1, 2, 3}, i: 2, want: []int{1, 2}, wantErr: nil},
		{name: "single_element_slice", xs: []int{42}, i: 0, want: []int{}, wantErr: nil},
		{name: "negative_index", xs: []int{1, 2, 3}, i: -1, want: nil, wantErr: sliceutil.ErrIndexOutOfRange},
		{name: "index_too_large", xs: []int{1, 2, 3}, i: 3, want: nil, wantErr: sliceutil.ErrIndexOutOfRange},
		{name: "empty_slice", xs: []int{}, i: 0, want: nil, wantErr: sliceutil.ErrIndexOutOfRange},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sliceutil.RemoveUnordered(tt.xs, tt.i)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("RemoveUnordered() error = %v, want %v", err, tt.wantErr)
			}
			if err == nil {
				if len(got) != len(tt.want) {
					t.Fatalf("RemoveUnordered() length = %v, want %v", len(got), len(tt.want))
				}
				for j := range got {
					if got[j] != tt.want[j] {
						t.Errorf("RemoveUnordered()[%d] = %v, want %v", j, got[j], tt.want[j])
					}
				}
			}
		})
	}
}

func TestPop(t *testing.T) {
	type test[T any] struct {
		name    string
		xs      []T
		wantXs  []T
		wantVal T
		wantErr error
	}

	tests := []test[int]{
		{
			name:    "pop_from_multiple",
			xs:      []int{1, 2, 3},
			wantXs:  []int{1, 2},
			wantVal: 3,
			wantErr: nil,
		},
		{
			name:    "pop_single_element",
			xs:      []int{42},
			wantXs:  []int{},
			wantVal: 42,
			wantErr: nil,
		},
		{
			name:    "pop_empty_slice",
			xs:      []int{},
			wantXs:  []int{},
			wantVal: 0,
			wantErr: sliceutil.ErrIndexOutOfRange,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotXs, gotVal, err := sliceutil.Pop(tt.xs)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("Pop() error = %v, want %v", err, tt.wantErr)
			}

			if err == nil {
				if len(gotXs) != len(tt.wantXs) {
					t.Fatalf("Pop() slice length = %v, want %v", len(gotXs), len(tt.wantXs))
				}
				for i := range gotXs {
					if gotXs[i] != tt.wantXs[i] {
						t.Errorf("Pop()[%d] = %v, want %v", i, gotXs[i], tt.wantXs[i])
					}
				}
				if gotVal != tt.wantVal {
					t.Errorf("Pop() value = %v, want %v", gotVal, tt.wantVal)
				}
			}
		})
	}
}

func TestRemoveLast(t *testing.T) {
	type test[T any] struct {
		name    string
		xs      []T
		want    []T
		wantErr error
	}

	tests := []test[int]{
		{name: "remove_last_from_multiple", xs: []int{1, 2, 3}, want: []int{1, 2}, wantErr: nil},
		{name: "remove_last_single_element", xs: []int{42}, want: []int{}, wantErr: nil},
		{name: "remove_last_empty_slice", xs: []int{}, want: nil, wantErr: sliceutil.ErrIndexOutOfRange},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sliceutil.RemoveLast(tt.xs)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("RemoveLast() error = %v, want %v", err, tt.wantErr)
			}
			if err == nil {
				if len(got) != len(tt.want) {
					t.Fatalf("RemoveLast() length = %v, want %v", len(got), len(tt.want))
				}
				for i := range got {
					if got[i] != tt.want[i] {
						t.Errorf("RemoveLast()[%d] = %v, want %v", i, got[i], tt.want[i])
					}
				}
			}
		})
	}
}
