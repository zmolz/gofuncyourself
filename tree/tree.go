package tree

import "cmp"

type Tree[K cmp.Ordered, V any] interface {
	Insert(k K, v V)
	Delete(k K) bool
	Get(k K) (V, bool)

	Size() int

	// Range
	Max() (K, V, bool)
	Min() (K, V, bool)

	String() string
}