package sliceutil_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/zmolz/gofuncyourself/sliceutil"
)

func TestTake(t *testing.T) {
	xs := []int{1, 2, 3, 4, 5}

	t.Run("take_partial", func(t *testing.T) {
		got, err := sliceutil.Take(xs, 3)
		want := []int{1, 2, 3}

		if err != nil {
			t.Fatalf("Take returned unexpected error: %v", err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Take(xs, 3) = %v, want %v", got, want)
		}
	})

	t.Run("take_all", func(t *testing.T) {
		got, err := sliceutil.Take(xs, 5)
		want := []int{1, 2, 3, 4, 5}

		if err != nil {
			t.Fatalf("Take returned unexpected error: %v", err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Take(xs, 5) = %v, want %v", got, want)
		}
	})

	t.Run("take_zero", func(t *testing.T) {
		got, err := sliceutil.Take(xs, 0)
		want := []int{}

		if err != nil {
			t.Fatalf("Take returned unexpected error: %v", err)
		}
		if len(got) != 0 {
			t.Errorf("Take(xs, 0) = %v, want empty slice", got)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Take(xs, 0) = %v, want %v", got, want)
		}
	})

	t.Run("take_too_many", func(t *testing.T) {
		got, err := sliceutil.Take(xs, 10)
		if !errors.Is(err, sliceutil.ErrIndexOutOfRange) {
			t.Errorf("Take(xs, 10) error = %v, want ErrIndexOutOfRange", err)
		}
		if got != nil {
			t.Errorf("Take(xs, 10) = %v, want nil slice", got)
		}
	})
}

func TestDrop(t *testing.T) {
	xs := []int{1, 2, 3, 4, 5}

	t.Run("drop_partial", func(t *testing.T) {
		got, err := sliceutil.Drop(xs, 2)
		want := []int{3, 4, 5}

		if err != nil {
			t.Fatalf("Drop returned unexpected error: %v", err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Drop(xs, 2) = %v, want %v", got, want)
		}
	})

	t.Run("drop_all", func(t *testing.T) {
		got, err := sliceutil.Drop(xs, 5)
		want := []int{}

		if err != nil {
			t.Fatalf("Drop returned unexpected error: %v", err)
		}
		if len(got) != 0 {
			t.Errorf("Drop(xs, 5) = %v, want empty slice", got)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Drop(xs, 5) = %v, want %v", got, want)
		}
	})

	t.Run("drop_zero", func(t *testing.T) {
		got, err := sliceutil.Drop(xs, 0)
		want := []int{1, 2, 3, 4, 5}

		if err != nil {
			t.Fatalf("Drop returned unexpected error: %v", err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Drop(xs, 0) = %v, want %v", got, want)
		}
	})

	t.Run("drop_too_many", func(t *testing.T) {
		got, err := sliceutil.Drop(xs, 10)
		if !errors.Is(err, sliceutil.ErrIndexOutOfRange) {
			t.Errorf("Drop(xs, 10) error = %v, want ErrIndexOutOfRange", err)
		}
		if got != nil {
			t.Errorf("Drop(xs, 10) = %v, want nil slice", got)
		}
	})
}
