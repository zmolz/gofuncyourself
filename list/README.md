``` go
type SkipList[K cmp.Ordered, V any] struct {
	// fields are not necessary to be shown
}

type SkipListNode[K cmp.Ordered, V any] struct {
	key     K
	value   V
	forward []*SkipListNode[K, V]
	level   int
}

func NewSkipList[K cmp.Ordered, V any]() *SkipList[K, V]

func (sl *SkipList[K, V]) Insert(k K, v V) 
func (sl *SkipList[K, V]) String() string

```