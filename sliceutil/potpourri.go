package sliceutil

// FilterMap applies function f over each element of slice xs,
// excluding elements for which f returns false, and returns the result.
func FilterMap[T, U any](xs []T, f func(T) (U, bool)) []U {
	ys := []U{}

	for _, x := range xs {
		if y, ok := f(x); ok {
			ys = append(ys, y)
		}
	}

	return ys
}

// FlatMap applies a function f over each element of slice xs,
// and then flattens the resulting slice.
func FlatMap[T, U any](xs []T, f func(T) []U) []U {
	// return Flatten(Map(xs, f))

	flatYs := []U{}

	for _, x := range xs {
		flatYs = append(flatYs, f(x)...)
	}

	return flatYs
}
