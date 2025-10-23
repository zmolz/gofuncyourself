package sliceutil

// Map applies function f over each element of slice xs, and returns the result.
func Map[T, U any](xs []T, f func(T) U) []U {
	ys := make([]U, len(xs))

	for i, x := range xs {
		ys[i] = f(x)
	}

	return ys
}
