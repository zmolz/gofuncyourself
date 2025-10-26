package iterator

type MapIterator[T, U any] struct {
	src Iterator[U]
	f   func(U) T
}

// NewMapIterator creates a lazy MapIterator from src Iterator
func NewMapIterator[T, U any](src Iterator[U], f func(U) T) *MapIterator[T, U] {
	return &MapIterator[T, U]{
		src: src,
		f:   f,
	}
}

func (m *MapIterator[T, U]) HasNext() bool {
	return m.src.HasNext()
}

func (m *MapIterator[T, U]) Next() (val T, ok bool) {
	if !m.src.HasNext() {
		return val, false
	}

	srcVal, _ := m.src.Next()

	return m.f(srcVal), true
}

