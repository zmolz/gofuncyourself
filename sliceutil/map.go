package sliceutil

// Map applies function f over each element of slice xs, and returns the result.
func Map[T any, U any](xs []T, f func(T) U) []U {
	ret := make([]U, len(xs))

	for i, x := range xs {
		ret[i] = f(x)
	}

	return ret
}
