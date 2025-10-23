package sliceutil_test

import (
	"slices"
	"testing"

	"github.com/zmolz/gofuncyourself/sliceutil"
)

func TestCompose(t *testing.T) {
	t.Run("int_to_string_then_len", func(t *testing.T) {
		f := func(x int) string { return string(rune('a' + x)) }
		g := func(s string) int { return len(s) }

		h := sliceutil.Compose(f, g)

		got := h(2)
		want := 1

		if got != want {
			t.Errorf("Compose(int->string->int)(2) = %v, want %v", got, want)
		}
	})

	t.Run("square_then_double", func(t *testing.T) {
		f := func(x int) int { return x * x }
		g := func(y int) int { return y * 2 }

		h := sliceutil.Compose(f, g)

		got := h(3)
		want := 18 // (3²)*2

		if got != want {
			t.Errorf("Compose(square->double)(3) = %v, want %v", got, want)
		}
	})

	t.Run("string_to_len_then_even_check", func(t *testing.T) {
		f := func(s string) int { return len(s) }
		g := func(n int) bool { return n%2 == 0 }

		h := sliceutil.Compose(f, g)

		if got := h("abcd"); !got {
			t.Errorf("Compose(string->len->even)(abcd) = false, want true")
		}
		if got := h("abc"); got {
			t.Errorf("Compose(string->len->even)(abc) = true, want false")
		}
	})

	t.Run("identity_composition", func(t *testing.T) {
		f := func(x int) int { return x }
		g := func(y int) int { return y }

		h := sliceutil.Compose(f, g)

		got := h(42)
		want := 42

		if got != want {
			t.Errorf("Compose(identity, identity)(42) = %v, want %v", got, want)
		}
	})

	t.Run("flat_map_composition", func(t *testing.T) {
		posIntLessThan := func(x int) []int {
			ret := make([]int, x)
			for i := range x {
				ret[i] = i + 1
			}
			return ret
		}
		f := func(xs []int) [][]int {
			return sliceutil.Map(xs, posIntLessThan)
		}
		g := sliceutil.Flatten[int]

		h := sliceutil.Compose(f, g)

		got := h([]int{1, 2, 3})
		want := sliceutil.FlatMap([]int{1, 2, 3}, posIntLessThan)

		if !slices.Equal(got, want) {
			t.Errorf("Compose(identity, identity)(42) = %v, want %v", got, want)
		}
	})
}
