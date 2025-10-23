package sliceutil_test

import (
	"testing"

	"github.com/zmolz/gofuncyourself/sliceutil"
)

func TestFind(t *testing.T) {
	xs := []int{1, 2, 3, 4, 5}

	t.Run("found", func(t *testing.T) {
		got, ok := sliceutil.Find(xs, func(x int) bool { return x > 3 })
		want := 4

		if !ok || got != want {
			t.Errorf("Find(xs, x>3) = (%v, %v), want (%v, true)", got, ok, want)
		}
	})

	t.Run("not_found", func(t *testing.T) {
		got, ok := sliceutil.Find(xs, func(x int) bool { return x > 10 })
		if ok {
			t.Errorf("Find(xs, x>10) = (%v, %v), want (_, false)", got, ok)
		}
	})
}

func TestFindAll(t *testing.T) {
	xs := []int{1, 2, 3, 4, 5, 6}
	got := sliceutil.FindAll(xs, func(x int) bool { return x%2 == 0 })
	want := []int{2, 4, 6}

	if len(got) != len(want) {
		t.Fatalf("FindAll returned wrong length: got %v, want %v", got, want)
	}
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("FindAll(xs, even)[%d] = %v, want %v", i, got[i], want[i])
		}
	}
}

func TestFindIndex(t *testing.T) {
	xs := []string{"a", "b", "c", "d"}

	t.Run("found", func(t *testing.T) {
		idx, ok := sliceutil.FindIndex(xs, func(s string) bool { return s == "c" })
		if !ok || idx != 2 {
			t.Errorf("FindIndex(xs, ==c) = (%v, %v), want (2, true)", idx, ok)
		}
	})

	t.Run("not_found", func(t *testing.T) {
		idx, ok := sliceutil.FindIndex(xs, func(s string) bool { return s == "z" })
		if ok {
			t.Errorf("FindIndex(xs, ==z) = (%v, %v), want (_, false)", idx, ok)
		}
	})
}


func TestFindIndices(t *testing.T) {
	xs := []int{1, 2, 3, 4, 5, 6}
	got := sliceutil.FindIndices(xs, func(x int) bool { return x%2 == 0 })
	want := []int{1, 3, 5}

	if len(got) != len(want) {
		t.Fatalf("FindIndices returned wrong length: got %v, want %v", got, want)
	}
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("FindIndices(xs, even)[%d] = %v, want %v", i, got[i], want[i])
		}
	}
}

func TestContains(t *testing.T) {
	xs := []int{10, 20, 30}

	if !sliceutil.Contains(xs, 20) {
		t.Errorf("Contains(xs, 20) = false, want true")
	}
	if sliceutil.Contains(xs, 40) {
		t.Errorf("Contains(xs, 40) = true, want false")
	}
}

func TestAny(t *testing.T) {
	xs := []int{1, 3, 5, 8}

	if !sliceutil.Any(xs, func(x int) bool { return x%2 == 0 }) {
		t.Errorf("Any(xs, even) = false, want true")
	}
	if sliceutil.Any(xs, func(x int) bool { return x > 10 }) {
		t.Errorf("Any(xs, >10) = true, want false")
	}
}

func TestAll(t *testing.T) {
	xs := []int{2, 4, 6}

	if !sliceutil.All(xs, func(x int) bool { return x%2 == 0 }) {
		t.Errorf("All(xs, even) = false, want true")
	}

	if sliceutil.All(xs, func(x int) bool { return x > 3 }) {
		t.Errorf("All(xs, >3) = true, want false")
	}
}
