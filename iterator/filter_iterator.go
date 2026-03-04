package iterator

type FilterIterator[T any] struct {
	src     Iterator[T]
	f       func(T) bool
	nextVal T
	nextOk  bool
	looked  bool
}

// NewFilterIterator creates a lazy FilterIterator from src Iterator
func NewFilterIterator[T any](src Iterator[T], f func(T) bool) *FilterIterator[T] {
	return &FilterIterator[T]{
		src: src,
		f:   f,
	}
}

func (filter *FilterIterator[T]) HasNext() bool {
	if filter.looked {
		return filter.nextOk
	}

	for filter.src.HasNext() {
		val, _ := filter.src.Next()
		if filter.f(val) {
			filter.nextVal = val
			filter.nextOk = true
			filter.looked = true
			return true
		}
	}

	filter.nextOk = false
	filter.looked = true
	return false
}

func (filter *FilterIterator[T]) Next() (T, bool) {
	if !filter.looked {
		if !filter.HasNext() {
			var zero T
			return zero, false
		}
	}

	filter.looked = false
	if filter.nextOk {
		return filter.nextVal, true
	}

	var zero T
	return zero, false
}
