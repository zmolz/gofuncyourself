package sliceutil_test

import (
	"reflect"
	"testing"

	"github.com/zmolz/gofuncyourself/sliceutil"
)

func TestZip(t *testing.T) {
	t.Run("equal length", func(t *testing.T) {
		xs := []int{1, 2, 3}
		ys := []string{"a", "b", "c"}

		got := sliceutil.Zip(xs, ys)
		want := []sliceutil.Pair[int, string]{
			{1, "a"},
			{2, "b"},
			{3, "c"},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Zip() = %v, want %v", got, want)
		}
	})

	t.Run("unequal length xs shorter", func(t *testing.T) {
		xs := []int{1, 2}
		ys := []string{"a", "b", "c"}

		got := sliceutil.Zip(xs, ys)
		want := []sliceutil.Pair[int, string]{
			{1, "a"},
			{2, "b"},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Zip() = %v, want %v", got, want)
		}
	})

	t.Run("unequal length ys shorter", func(t *testing.T) {
		xs := []int{1, 2, 3}
		ys := []string{"a"}

		got := sliceutil.Zip(xs, ys)
		want := []sliceutil.Pair[int, string]{
			{1, "a"},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Zip() = %v, want %v", got, want)
		}
	})

	t.Run("empty slices", func(t *testing.T) {
		var xs []int
		var ys []string
		got := sliceutil.Zip(xs, ys)
		if len(got) != 0 {
			t.Errorf("Zip() = %v, want empty slice", got)
		}
	})
}

func TestZipLongest(t *testing.T) {
	t.Run("equal length", func(t *testing.T) {
		xs := []int{1, 2, 3}
		ys := []string{"a", "b", "c"}

		got := sliceutil.ZipLongest(xs, ys)
		want := []sliceutil.Pair[int, string]{
			{1, "a"},
			{2, "b"},
			{3, "c"},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("ZipLongest() = %v, want %v", got, want)
		}
	})

	t.Run("xs shorter", func(t *testing.T) {
		xs := []int{1, 2}
		ys := []string{"a", "b", "c"}

		got := sliceutil.ZipLongest(xs, ys)
		want := []sliceutil.Pair[int, string]{
			{1, "a"},
			{2, "b"},
			{0, "c"}, // zero value of int for missing xs
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("ZipLongest() = %v, want %v", got, want)
		}
	})

	t.Run("ys shorter", func(t *testing.T) {
		xs := []int{1, 2, 3}
		ys := []string{"a"}

		got := sliceutil.ZipLongest(xs, ys)
		want := []sliceutil.Pair[int, string]{
			{1, "a"},
			{2, ""},
			{3, ""},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("ZipLongest() = %v, want %v", got, want)
		}
	})
}

func TestZipWith(t *testing.T) {
	t.Run("sum ints", func(t *testing.T) {
		xs := []int{1, 2, 3}
		ys := []int{4, 5, 6}

		got := sliceutil.ZipWith(xs, ys, func(a, b int) int { return a + b })
		want := []int{5, 7, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("ZipWith() = %v, want %v", got, want)
		}
	})

	t.Run("concatenate strings", func(t *testing.T) {
		xs := []string{"a", "b"}
		ys := []string{"x", "y", "z"}

		got := sliceutil.ZipWith(xs, ys, func(a, b string) string { return a + b })
		want := []string{"ax", "by"}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("ZipWith() = %v, want %v", got, want)
		}
	})
}

func TestUnzip(t *testing.T) {
	t.Run("basic case", func(t *testing.T) {
		pairs := []sliceutil.Pair[int, string]{
			{1, "a"},
			{2, "b"},
			{3, "c"},
		}

		gotX, gotY := sliceutil.Unzip(pairs)
		wantX := []int{1, 2, 3}
		wantY := []string{"a", "b", "c"}

		if !reflect.DeepEqual(gotX, wantX) {
			t.Errorf("Unzip() xs = %v, want %v", gotX, wantX)
		}
		if !reflect.DeepEqual(gotY, wantY) {
			t.Errorf("Unzip() ys = %v, want %v", gotY, wantY)
		}
	})

	t.Run("empty", func(t *testing.T) {
		var pairs []sliceutil.Pair[int, string]
		gotX, gotY := sliceutil.Unzip(pairs)
		if len(gotX) != 0 || len(gotY) != 0 {
			t.Errorf("Unzip() empty = (%v, %v), want empty slices", gotX, gotY)
		}
	})
}
