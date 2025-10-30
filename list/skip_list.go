package list

import (
	"cmp"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	maxLevel = 32
	pPercent = 25
	pScaled  = uint32(pPercent * 65536 / 100)
)

/*
optimizations
	backwards pointers at level 0 -- easy access to max
	current max level

features
	span counters for nodes, for O(log n) "indexing" / rank queries
*/

type SkipList[K cmp.Ordered, V any] struct {
	header   *SkipListNode[K, V]
	sentinel *SkipListNode[K, V]

	// having one preallocated array is better than having to allocate one every time
	// i go to insert or delete, so im making this a field of the skiplist.
	updates [maxLevel]*SkipListNode[K, V]

	length int

	r *rand.Rand
}

type SkipListNode[K cmp.Ordered, V any] struct {
	key     K
	value   V
	forward []*SkipListNode[K, V]
	level   int
}

func NewSkipList[K cmp.Ordered, V any]() *SkipList[K, V] {
	sentinel := &SkipListNode[K, V]{}

	forward := make([]*SkipListNode[K, V], maxLevel)
	for i := range maxLevel {
		forward[i] = sentinel
	}
	header := &SkipListNode[K, V]{forward: forward, level: maxLevel}

	return &SkipList[K, V]{
		header:   header,
		sentinel: sentinel,

		length: 0,

		r: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (sl *SkipList[K, V]) Insert(k K, v V) {
	curr := sl.header

	// at the end of this loop, curr will be the node immediately prior to the value to add
	for level := maxLevel - 1; level >= 0; level-- {
		// iterate across at current level
		// until we either hit the sentinel or a key > the key to insert
		// in which case we drop a down a level.
		for curr.forward[level] != sl.sentinel && curr.forward[level].key < k {
			curr = curr.forward[level]
		}
		sl.updates[level] = curr // record last node before dropping a level
	}

	// update node?
	next := curr.forward[0]
	if next != sl.sentinel && next.key == k {
		next.value = v
		return
	}

	// add new node
	newLevel := randomLevel(sl.r)
	newNode := &SkipListNode[K, V]{
		key:     k,
		value:   v,
		forward: make([]*SkipListNode[K, V], newLevel),
		level:   newLevel,
	}
	sl.length++

	// update the forwards of prior nodes, and set the forwards of current node
	for level := range newLevel {
		newNode.forward[level] = sl.updates[level].forward[level]
		sl.updates[level].forward[level] = newNode
	}
}

func randomLevel(r *rand.Rand) int {
	level := 1
	for (r.Uint32()&0xFFFF) < uint32(pScaled) && level < maxLevel {
		level++
	}
	return level
}

func (sl *SkipList[K, V]) Delete(k K) bool {
	curr := sl.header

	// iterate across at a level, until we hit a key > the key to delete or the sentinel
	// in which case we move down, then repeat until curr.forward[0] == potential node to delete
	for level := maxLevel - 1; level >= 0; level-- {
		for curr.forward[level] != sl.sentinel && curr.forward[level].key < k {
			curr = curr.forward[level]
		}
		sl.updates[level] = curr
	}

	target := curr.forward[0]
	if target == sl.sentinel || target.key != k {
		return false
	}

	// start from 0 to the height of the node to unlink,
	// link prior node on that level to the next node on that level
	// if target is between them
	for level := range target.level {
		if sl.updates[level].forward[level] == target {
			sl.updates[level].forward[level] = target.forward[level]
		}
	}

	sl.length--
	return true
}

func (sl *SkipList[K, V]) Length() int {
	return sl.length
}

func (sl *SkipList[K, V]) Get(k K) (V, bool) {
	curr := sl.header

	// you get the gist by now
	for level := maxLevel - 1; level >= 0; level-- {
		for curr.forward[level] != sl.sentinel && curr.forward[level].key < k {
			curr = curr.forward[level]
		}
	}

	// order up
	next := curr.forward[0]
	if next.key == k {
		return next.value, true
	}

	var zero V
	return zero, false
}

func (sl *SkipList[K, V]) Range(lo, hi K) []V {
	curr := sl.header

	// search for first node before range
	for level := maxLevel - 1; level >= 0; level-- {
		for curr.forward[level] != sl.sentinel && curr.forward[level].key < lo {
			curr = curr.forward[level]
		}
	}

	// the next node, if it is not the sentinel, will be the first node > lo
	// aka potentially the first node in the range if it is not > hi.
	// just iterate across bottom layer of list.
	values := []V{}
	// inclusive range
	for curr := curr.forward[0]; curr != sl.sentinel && curr.key <= hi; curr = curr.forward[0] {
		values = append(values, curr.value)
	}

	return values
}

// once we add backwards pointers at level 0 this will be easy
func (sl *SkipList[K, V]) Max() (k K, v V, ok bool) {
	if sl.length == 0 {
		return
	}

	curr := sl.header

	for level := maxLevel - 1; level >= 0; level-- {
		for curr.forward[level] != sl.sentinel {
			curr = curr.forward[level]
		}
	}

	return curr.key, curr.value, true
}

func (sl *SkipList[K, V]) Min() (k K, v V, ok bool) {
	if sl.length == 0 {
		return
	}

	minNode := sl.header.forward[0]
	return minNode.key, minNode.value, true
}

// this was easier than i expected
func (sl *SkipList[K, V]) String() string {
	repr := strings.Builder{}

	// iterate through bottom, once, collecting nodes in order.
	nodes := make([]*SkipListNode[K, V], sl.length)
	curr := sl.header
	for i := range sl.length {
		curr = curr.forward[0]
		nodes[i] = curr
	}

	// print all levels (top to bottom), showing pointers
	for level := maxLevel - 1; level >= 0; level-- {

		// start with header
		repr.WriteString("\n⎪    ⎪")

		for i, node := range nodes {

			if node.level <= level {
				repr.WriteString("---------")
			} else {

				// gonna have to fix this later to be more generalized
				repr.WriteString(fmt.Sprintf("-▶⎪%5v⎪", node.value))
			}

			if i+1 == len(nodes) {
				// next is sentinel, draw arrow
				repr.WriteString("-▶")
			} else {
				// next is not sentinel
				repr.WriteString("--")
			}
		}

		// finish line with sentinel
		repr.WriteString("⎪    ⎪")
	}

	repr.WriteString("\n")

	return repr.String()
}
