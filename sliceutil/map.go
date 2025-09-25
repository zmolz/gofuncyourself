package sliceutil

func Map[T any, U any](f func(T) U, xs []T) []U {
	ret := make([]U, len(xs));

	for i, x := range xs {
		ret[i]= f(x)
	}

	return ret
}