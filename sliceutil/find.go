package sliceutil

import "cmp"

// Find returns the first element of slice xs where f(x) is true
func Find[T any](xs []T, f func(T) bool) (T, bool) {
	for _, x := range xs {
		if f(x) {
			return x, true
		}
	}

	var zero T
	return zero, false
}

// FindAll returns all elements of slice xs where f(x) is true
func FindAll[T any](xs []T, f func(T) bool) []T {
	found := []T{}

	for _, x := range xs {
		if f(x) {
			found = append(found, x)
		}
	}

	return found
}

// FindIndex returns the index in slice xs of the first element where f(x) is true
func FindIndex[T any](xs []T, f func(T) bool) (int, bool) {
	for i, x := range xs {
		if f(x) {
			return i, true
		}
	}

	var zero int
	return zero, false
}

// FindIndices returns all indices in slice xs of elemente where f(x) is true
func FindIndices[T any](xs []T, f func(T) bool) []int {
	indices := []int{}

	for i, x := range xs {
		if f(x) {
			indices = append(indices, i)
		}
	}

	return indices
}

// Contains returns true if slice xs contains an element
func Contains[T cmp.Ordered](xs []T, toFind T) bool {
	for _, x := range xs {
		if x == toFind {
			return true
		}
	}

	return false
}

// Any returns true if any element of slice xs returns true for function f
func Any[T any](xs []T, f func(T) bool) bool {
	for _, x := range xs {
		if f(x) {
			return true
		}
	}

	return false
}

// All returns true if all elements of slice xs return true for function f
func All[T any](xs []T, f func(T) bool) bool {
	for _, x := range xs {
		if !f(x) {
			return false
		}
	}

	return true
}
