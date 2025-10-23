package sliceutil

// Filter returns all elements of slice xs for which predicate function pred is true
func Filter[T any](xs []T, pred func(T) bool) []T {
	filtered := []T{}

	for _, x := range xs {
		if pred(x) {
			filtered = append(filtered, x)
		}
	}

	return filtered
}

// Partition split slice xs into two slices,
// the first being all elements of xs for which predicate function pred is true,
// the second being all elements for which pred is false.
func Partition[T any](xs []T, pred func(T) bool) (trueElements []T, falseElements []T) {
	for _, x := range xs {
		if pred(x) {
			trueElements = append(trueElements, x)
		} else {
			falseElements = append(falseElements, x)
		}
	}

	return trueElements, falseElements
}

// GroupBy returns a map such that each element of slice xs exists in the map
// as a value in a list, keyed by the given function keyFunc. 
func GroupBy[T any, U comparable](xs []T, keyFunc func(T) U) map[U][]T {
	groups := map[U][]T{}

	for _, x := range xs {
		key := keyFunc(x)
		groups[key] = append(groups[key], x)
	}

	return groups
}