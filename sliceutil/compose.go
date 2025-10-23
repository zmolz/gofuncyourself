package sliceutil

// Compose is... pretty useless in Go but here it is.
// Return a function that applies function f, then function g to a value.
func Compose[T, U, V any](f func(T) U, g func(U) V) func(T) V {
	return func(t T) V {
		return g(f(t))
	}
}