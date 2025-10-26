package tree

import "cmp"

// BST represents a binary search tree
type BST[K cmp.Ordered, V any] struct {
	Root *Node[K, V]
}

// Insert a key-value pair into the tree
func (t *BST[K, V]) Insert(k K, v V) {
	t.Root = insert(t.Root, k, v)
}

// Delete a key from the tree
func (t *BST[K, V]) Delete(k K) bool {
	root, deleted := delete(t.Root, k)
	t.Root = root
	return deleted
}

// Get looks up a key in the BST. It returns the value and true if found,
// or the zero value and false if the key is not present
func (t *BST[K, V]) Get(k K) (V, bool) {
	return get(t.Root, k)
}

// Size returns the amount of nodes in the tree
func (t *BST[K, V]) Size() int {
	return size(t.Root)
}

// Range returns all V values who's keys are between K l and r.
func (t *BST[K, V]) Range(l, r K) []V {
	return getWithinRange(t.Root, l, r)
}

// Max returns the key and value associated with the maximum key in the tree
func (t *BST[K, V]) Max() (maxKey K, maxVal V, ok bool) {
	if t.Root == nil {
		return maxKey, maxVal, false
	}
	maxNode := max(t.Root)
	return maxNode.Key, maxNode.Value, true
}

// Min returns the key and value associated with the minimum key in the tree
func (t *BST[K, V]) Min() (minKey K, minVal V, ok bool) {
	if t.Root == nil {
		return minKey, minVal, false
	}
	minNode := min(t.Root)
	return minNode.Key, minNode.Value, true
}

// DebugPrint prints the tree structure to the console for debugging.
func (t *BST[K, V]) String() string {
	if t.Root == nil {
		return "<empty tree>"
	}
	return stringRepr(t.Root, 0)
}
