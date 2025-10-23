package sliceutil_test

import (
	"testing"

	"github.com/zmolz/gofuncyourself/sliceutil"
)

func TestFilterInt(t *testing.T) {
	type test struct {
		name string
		xs   []int
		pred func(int) bool
		want []int
	}

	isEven := func(x int) bool { return x%2 == 0 }
	isOdd := func(x int) bool { return x%2 != 0 }
	alwaysTrue := func(x int) bool { return true }
	alwaysFalse := func(x int) bool { return false }

	tests := []test{
		{
			name: "filter_even_numbers",
			xs:   []int{1, 2, 3, 4, 5, 6},
			pred: isEven,
			want: []int{2, 4, 6},
		},
		{
			name: "filter_odd_numbers",
			xs:   []int{1, 2, 3, 4, 5, 6},
			pred: isOdd,
			want: []int{1, 3, 5},
		},
		{
			name: "filter_all_true",
			xs:   []int{1, 2, 3},
			pred: alwaysTrue,
			want: []int{1, 2, 3},
		},
		{
			name: "filter_all_false",
			xs:   []int{1, 2, 3},
			pred: alwaysFalse,
			want: []int{},
		},
		{
			name: "filter_empty_slice",
			xs:   []int{},
			pred: isEven,
			want: []int{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := sliceutil.Filter(test.xs, test.pred)
			if len(got) != len(test.want) {
				t.Fatalf("Filter() length = %v, want %v", len(got), len(test.want))
			}
			for i := range got {
				if got[i] != test.want[i] {
					t.Errorf("Filter()[%d] = %v, want %v", i, got[i], test.want[i])
				}
			}
		})
	}
}

func TestFilterString(t *testing.T) {
	type test struct {
		name string
		xs   []string
		pred func(string) bool
		want []string
	}

	isLong := func(s string) bool { return len(s) > 3 }

	tests := []test{
		{
			name: "filter_long_strings",
			xs:   []string{"go", "gopher", "GPT", "chatGPT"},
			pred: isLong,
			want: []string{"gopher", "chatGPT"},
		},
		{
			name: "filter_empty_slice",
			xs:   []string{},
			pred: isLong,
			want: []string{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := sliceutil.Filter(test.xs, test.pred)
			if len(got) != len(test.want) {
				t.Fatalf("Filter() length = %v, want %v", len(got), len(test.want))
			}
			for i := range got {
				if got[i] != test.want[i] {
					t.Errorf("Filter()[%d] = %v, want %v", i, got[i], test.want[i])
				}
			}
		})
	}
}

func TestPartitionInt(t *testing.T) {
	type test struct {
		name      string
		xs        []int
		pred      func(int) bool
		wantTrue  []int
		wantFalse []int
	}

	isEven := func(x int) bool { return x%2 == 0 }
	alwaysTrue := func(x int) bool { return true }
	alwaysFalse := func(x int) bool { return false }

	tests := []test{
		{
			name:      "partition_even_numbers",
			xs:        []int{1, 2, 3, 4, 5, 6},
			pred:      isEven,
			wantTrue:  []int{2, 4, 6},
			wantFalse: []int{1, 3, 5},
		},
		{
			name:      "partition_all_true",
			xs:        []int{1, 2, 3},
			pred:      alwaysTrue,
			wantTrue:  []int{1, 2, 3},
			wantFalse: []int{},
		},
		{
			name:      "partition_all_false",
			xs:        []int{1, 2, 3},
			pred:      alwaysFalse,
			wantTrue:  []int{},
			wantFalse: []int{1, 2, 3},
		},
		{
			name:      "partition_empty_slice",
			xs:        []int{},
			pred:      isEven,
			wantTrue:  []int{},
			wantFalse: []int{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotTrue, gotFalse := sliceutil.Partition(test.xs, test.pred)

			if len(gotTrue) != len(test.wantTrue) {
				t.Fatalf("Partition() true slice length = %v, want %v", len(gotTrue), len(test.wantTrue))
			}
			for i := range gotTrue {
				if gotTrue[i] != test.wantTrue[i] {
					t.Errorf("Partition() true slice[%d] = %v, want %v", i, gotTrue[i], test.wantTrue[i])
				}
			}

			if len(gotFalse) != len(test.wantFalse) {
				t.Fatalf("Partition() false slice length = %v, want %v", len(gotFalse), len(test.wantFalse))
			}
			for i := range gotFalse {
				if gotFalse[i] != test.wantFalse[i] {
					t.Errorf("Partition() false slice[%d] = %v, want %v", i, gotFalse[i], test.wantFalse[i])
				}
			}
		})
	}
}

func TestPartitionString(t *testing.T) {
	type test struct {
		name      string
		xs        []string
		pred      func(string) bool
		wantTrue  []string
		wantFalse []string
	}

	isLong := func(s string) bool { return len(s) > 3 }

	tests := []test{
		{
			name:      "partition_long_strings",
			xs:        []string{"go", "gopher", "GPT", "chatGPT"},
			pred:      isLong,
			wantTrue:  []string{"gopher", "chatGPT"},
			wantFalse: []string{"go", "GPT"},
		},
		{
			name:      "partition_empty_slice",
			xs:        []string{},
			pred:      isLong,
			wantTrue:  []string{},
			wantFalse: []string{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotTrue, gotFalse := sliceutil.Partition(test.xs, test.pred)

			if len(gotTrue) != len(test.wantTrue) {
				t.Fatalf("Partition() true slice length = %v, want %v", len(gotTrue), len(test.wantTrue))
			}
			for i := range gotTrue {
				if gotTrue[i] != test.wantTrue[i] {
					t.Errorf("Partition() true slice[%d] = %v, want %v", i, gotTrue[i], test.wantTrue[i])
				}
			}

			if len(gotFalse) != len(test.wantFalse) {
				t.Fatalf("Partition() false slice length = %v, want %v", len(gotFalse), len(test.wantFalse))
			}
			for i := range gotFalse {
				if gotFalse[i] != test.wantFalse[i] {
					t.Errorf("Partition() false slice[%d] = %v, want %v", i, gotFalse[i], test.wantFalse[i])
				}
			}
		})
	}
}

func TestGroupBy(t *testing.T) {
	t.Run("group_by_first_letter", func(t *testing.T) {
		firstByte := func(s string) byte {
			return s[0]
		}

		xs := []string{"ale", "armor", "bobcat", "cat", "christmas"}

		got := sliceutil.GroupBy(xs, firstByte)
		want := map[byte][]string{
			'a': {"ale", "armor"},
			'b': {"bobcat"},
			'c': {"cat", "christmas"},
		}

		if len(got) != len(want) {
			t.Fatalf("GroupBy length mismatch: got %v, want %v", got, want)
		}

		for key, wantVals := range want {
			gotVals, ok := got[key]
			if !ok {
				t.Errorf("GroupBy missing key %q", key)
				continue
			}

			if len(gotVals) != len(wantVals) {
				t.Errorf("GroupBy[%q] length = %d, want %d", key, len(gotVals), len(wantVals))
				continue
			}

			for i := range wantVals {
				if gotVals[i] != wantVals[i] {
					t.Errorf("GroupBy[%q][%d] = %v, want %v", key, i, gotVals[i], wantVals[i])
				}
			}
		}
	})
}
