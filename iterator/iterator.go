package iterator

type Iterator[T any] interface {
	// Increment iterator. If unexhausted, return next value and true, otherwise return zero value and false
	Next() (T, bool)

	// Return if iterator is unexhausted
	HasNext() bool
}

// Exhaust returns a slice of all remaining values in Iterator it
func Exhaust[T any](it Iterator[T]) []T {
	values := []T{}
	for it.HasNext() {
		curr, _ := it.Next()
		values = append(values, curr)
	}
	return values
}

