package sliceutil

import (
	"slices"
)

// Remove removes the element at index i from slice xs.
func Remove[T any](xs []T, i int) ([]T, error) {
	if i < 0 || i >= len(xs) {
		// for potential use in loops we might want to return xs...
		// nil makes sense too.
		return nil, ErrIndexOutOfRange
	}

	return append(xs[:i], xs[i+1:]...), nil
}

// RemoveUnordered removes the element at index i from slice xs in O(1) time, but does not preserve order.
func RemoveUnordered[T any](xs []T, i int) ([]T, error) {
	if i < 0 || i >= len(xs) {
		return nil, ErrIndexOutOfRange
	}
	// swap with last element
	xs[i] = xs[len(xs)-1]

	return xs[:len(xs)-1], nil
}

// Pop removes the last element of slice xs, and returns it.
func Pop[T any](xs []T) ([]T, T, error) {
	if len(xs) == 0 {
		var zero T
		return xs, zero, ErrIndexOutOfRange // perhaps switch to ErrRemoveFromEmptyList?
	}

	return xs[:len(xs)-1], xs[len(xs)-1], nil
}

// RemoveLast removes the last element of slice xs.
func RemoveLast[T any](xs []T) ([]T, error) {
	xs, _, err := Pop(xs)

	return xs, err
}

// RemoveIndices takes a slice of ints indices and removes the associated elements in slice xs
func RemoveIndices[T any](xs []T, indices []int) ([]T, error) {
	var err error

	// Sort descending to remove from the back first
	slices.Sort(indices)
	slices.Reverse(indices)

	for _, i := range indices {
		xs, err = RemoveUnordered(xs, i)
		if err != nil {
			return nil, err
		}
	}

	return xs, nil
}