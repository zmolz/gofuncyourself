package list

import (
	"cmp"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const maxLevel = 8
const p = 0.5

type SkipList[K cmp.Ordered, V any] struct {
	header   *SkipListNode[K, V]
	sentinel *SkipListNode[K, V]

	// uhhh this is hacky maybe but im trying to make it more efficient,
	// having one preallocated array is better than having to allocate one every time
	// i go to insert, so im making this a field of the skiplist.
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
	// start from top level of header
	curr := sl.header

	// at the end of this loop, curr will be the node immediately prior to the value to add
	level := maxLevel - 1
	for level >= 0 {

		// move down a level if forward[level] is sentinel, or > the key we want to add
		if curr.forward[level] == sl.sentinel || curr.forward[level].key > k {
			sl.updates[level] = curr // only record update when we move down
			level--
			continue
		}

		// key already exists in list, update value at node, and return
		if curr.forward[level].key == k {
			curr.forward[level].value = v
			return
		}

		// move across to curr.forward[level], key > curr.forward[level].key
		curr = curr.forward[level]
	}

	// add new node
	sl.length++

	newLevel := randomLevel(sl.r)
	newNode := &SkipListNode[K, V]{
		key:     k,
		value:   v,
		forward: make([]*SkipListNode[K, V], newLevel),
		level:   newLevel,
	}

	// update the forwards of prior nodes, and set the forwards of current node
	for level := range newLevel {
		newNode.forward[level] = sl.updates[level].forward[level]
		sl.updates[level].forward[level] = newNode
	}
}

func randomLevel(r *rand.Rand) int {
	level := 1
	for r.Float64() < p && level < maxLevel {
		level++
	}
	return level
}

// this was easier than i expected
func (sl *SkipList[K, V]) Repr() string {
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
				repr.WriteString(fmt.Sprintf("-▶⎪%3v  ⎪", node.key))
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
