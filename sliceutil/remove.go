package sliceutil

func Remove[T any](s []T, i int) []T {
	return append(s[:i], s[i+1:]...)
}
