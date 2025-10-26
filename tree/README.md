```go
type Tree[K cmp.Ordered, V any] interface {
	Insert(k K, v V)
	Delete(k K) bool

	Size() int

	Get(k K) (V, bool)
	Range(l, r K) []V
	Max() (K, V, bool)
	Min() (K, V, bool)

	String() string
}
```
