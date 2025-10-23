package sliceutil

import "errors"

// FoldLeft applies function f to each element of slice xs and an accumulator a, 
// from left to right, returning the final accumulated value.
// The function f takes the current element and the accumulator, and returns the updated accumulator.
func FoldLeft[T, U any](xs []T, a U, f func(T, U) U) U {
	for _, x := range xs {
		a = f(x, a)
	}
	return a
}

var ErrEmptyReduce = errors.New("Reduce() of empty slice with no initial value")
// Reduce combines the elements of slice xs using function f, returning a single value.
// If slice xs is empty, it returns error ErrEmptyReduce
func Reduce[T any](xs []T, f func(T, T) T) (T, error) {
	if len(xs) == 0 {
		var zero T
		return zero, ErrEmptyReduce
	}

	return FoldLeft(xs[1:], xs[0], f), nil
}

// FoldRight applies function f to each element of slice xs and an accumulator a, 
// from right to left, returning the final accumulated value.
// The function f takes the current element and the accumulator, and returns the updated accumulator.
func FoldRight[T, U any](xs []T, a U, f func(T, U) U, debug bool) U {
	for i := len(xs) - 1; i >= 0; i-- {
		a = f(xs[i], a)
	}

	return a
}
