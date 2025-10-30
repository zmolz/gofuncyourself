package list_test

import (
	"fmt"
	"testing"

	. "github.com/zmolz/gofuncyourself/list"
)

func TestInsertAndGet(t *testing.T) {
	sl := NewSkipList[int, string]()
	count := 100

	for i := range count {
		sl.Insert(i, fmt.Sprintf("val-%d", i))
	}

	if sl.Length() != count {
		t.Fatalf("expected length %d, got %d, %s", count, sl.Length(), sl)
	}

	for i := range count {
		v, ok := sl.Get(i)
		if !ok {
			t.Fatalf("expected key %d to exist", i)
		}
		if v != fmt.Sprintf("val-%d", i) {
			t.Fatalf("expected val-%d, got %ss", i, v)
			// t.Fatalf("%s", sl)
		}
	}

	// overwrite a value
	sl.Insert(50, "updated")
	if sl.Length() != count {
		t.Fatalf("expected length %d, got %d, %s", count, sl.Length(), sl)
	}
	v, ok := sl.Get(50)
	if !ok || v != "updated" {
		t.Fatalf("expected overwrite to succeed, got (%v, %v)", v, ok)
	}
}

func TestDelete(t *testing.T) {
	sl := NewSkipList[int, string]()

	for i := range 10 {
		sl.Insert(i, fmt.Sprintf("v%d", i))
	}

	ok := sl.Delete(5)
	if !ok {
		t.Fatalf("expected delete(5) to return true")
	}
	if sl.Length() != 9 {
		t.Fatalf("expected length 9 after delete, got %d", sl.Length())
	}

	_, ok = sl.Get(5)
	if ok {
		t.Fatalf("expected key 5 to be gone")
	}

	// deleting non-existent key should return false
	if sl.Delete(99) {
		t.Fatalf("expected delete(99) to return false")
	}
}

func TestRange(t *testing.T) {
	sl := NewSkipList[int, string]()
	for i := range 10 {
		sl.Insert(i, fmt.Sprintf("v%d", i))
	}

	vals := sl.Range(3, 6)
	expected := []string{"v3", "v4", "v5", "v6"}
	if len(vals) != len(expected) {
		t.Fatalf("expected %d results, got %d", len(expected), len(vals))
	}
	for i, v := range vals {
		if v != expected[i] {
			t.Fatalf("expected %s at index %d, got %s", expected[i], i, v)
		}
	}
}

func TestMinMax(t *testing.T) {
	sl := NewSkipList[int, string]()
	if _, _, ok := sl.Min(); ok {
		t.Fatalf("expected Min() on empty list to return ok=false")
	}
	if _, _, ok := sl.Max(); ok {
		t.Fatalf("expected Max() on empty list to return ok=false")
	}

	for i := 1; i <= 5; i++ {
		sl.Insert(i, fmt.Sprintf("v%d", i))
	}

	minK, minV, ok := sl.Min()
	if !ok || minK != 1 || minV != "v1" {
		t.Fatalf("unexpected Min(): (%v,%v,%v)", minK, minV, ok)
	}

	maxK, maxV, ok := sl.Max()
	if !ok || maxK != 5 || maxV != "v5" {
		t.Fatalf("unexpected Max(): (%v,%v,%v)", maxK, maxV, ok)
	}
}

func TestStringRepresentation(t *testing.T) {
	sl := NewSkipList[int, string]()
	for i := 0; i < 5; i++ {
		sl.Insert(i, fmt.Sprintf("v%d", i))
	}

	out := sl.String()
	if len(out) == 0 {
		t.Fatalf("expected non-empty string representation")
	}
	t.Logf("\n%s", out)
}

func TestInsertDeleteStress(t *testing.T) {
	sl := NewSkipList[int, int]()
	n := 1000

	for i := range n {
		sl.Insert(i, i)
	}

	for i := range n {
		if !sl.Delete(i) {
			t.Fatalf("delete failed at %d", i)
		}
	}

	if sl.Length() != 0 {
		t.Fatalf("expected empty list after full deletion, got length=%d", sl.Length())
	}
}
