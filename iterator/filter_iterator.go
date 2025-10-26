package iterator

type FilterIterator[T any] struct {
	src Iterator[T]
	f   func(T) bool
}

// NewFilterIterator creates a lazy FilterIterator from src Iterator
func NewFilterIterator[T any](src Iterator[T], f func(T) bool) *FilterIterator[T] {
	return &FilterIterator[T]{
		src: src,
		f:   f,
	}
}

func (filter *FilterIterator[T]) HasNext() bool {
	return filter.src.HasNext()
}

func (filter *FilterIterator[T]) Next() (T, bool) {
	for filter.src.HasNext() {
		val, _ := filter.src.Next()
		if filter.f(val) {
			return val, true
		}
	}

	var zero T
	return zero, false
}
