package iterator

type SliceIterator[T any] struct {
	src []T
	i   int
}

// NewMapIterator creates a lazy MapIterator from src Iterator
func NewSliceIterator[T any](src []T) *SliceIterator[T] {
	return &SliceIterator[T]{
		src: src,
		i:   -1,
	}
}

func (it *SliceIterator[T]) HasNext() bool {
	return it.i < len(it.src)-1
}

func (it *SliceIterator[T]) Next() (next T, ok bool) {
	if !it.HasNext() {
		return next, false
	}

 	it.i += 1
	return it.src[it.i], true
}
