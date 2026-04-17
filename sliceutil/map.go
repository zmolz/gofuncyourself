package sliceutil

import "errors"

// Map applies function f over each element of slice xs, and returns the result.
func Map[T, U any](xs []T, f func(T) U) []U {
	ys := make([]U, len(xs))

	for i, x := range xs {
		ys[i] = f(x)
	}

	return ys
}

// MapWithErr applies function f over each element of slice xs, and returns the result - while allowing for errors
func MapWithErr[T, U any](xs []T, f func(T) (U, error)) ([]U, error) {
	ys := make([]U, len(xs))
	es := make([]error, len(xs))

	for i, x := range xs {
		ys[i], es[i] = f(x)
	}

	return ys, errors.Join(es...)
}
