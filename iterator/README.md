``` go
type Iterator[T any] interface {
	Next() (T, bool)
	HasNext() bool
}

func Exhaust[T any](it Iterator[T]) []T

type FilterIterator[T any] struct
type MapIterator[T, U any] struct

type NumRange[T Num] struct
```