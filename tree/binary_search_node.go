package tree

import (
	"cmp"
	"fmt"
	"strings"
)

// Node represents a Node in the BST
type Node[K cmp.Ordered, V any] struct {
	Key   K
	Value V
	Left  *Node[K, V]
	Right *Node[K, V]
}

// insert inserts (k, v) into the subtree rooted at n
func insert[K cmp.Ordered, V any](n *Node[K, V], k K, v V) *Node[K, V] {
	if n == nil {
		return &Node[K, V]{Key: k, Value: v}
	}

	if k < n.Key {
		n.Left = insert(n.Left, k, v)
	} else if k > n.Key {
		n.Right = insert(n.Right, k, v)
	} else {
		n.Value = v
	}

	return n
}

// delete deletes a key k from the subtree rooted at n
func delete[K cmp.Ordered, V any](n *Node[K, V], k K) (*Node[K, V], bool) {
	if n == nil {
		return nil, false
	}

	var found bool
	if k < n.Key {
		n.Left, found = delete(n.Left, k)
	} else if k > n.Key {
		n.Right, found = delete(n.Right, k)
	} else {
		// found node to delete
		found = true

		// we are a leaf
		if n.Left == nil && n.Right == nil {
			return nil, found

		}

		// one child
		if n.Right != nil && n.Left == nil {
			return n.Right, found
		}
		if n.Left != nil && n.Right == nil {
			return n.Left, found
		}

		// two children -- swap key-value pair with inorder successor, aka next largest key
		successor := min(n.Right)
		n.Key = successor.Key
		n.Value = successor.Value

		// once we swap value, delete that node from the right subtree
		n.Right, _ = delete(n.Right, successor.Key)
	}

	return n, found
}

// get retrieves the value associated with a given key in the subtree rooted at n
func get[K cmp.Ordered, V any](n *Node[K, V], k K) (V, bool) {
	if n == nil {
		var zero V
		return zero, false
	}

	if k < n.Key {
		return get(n.Left, k)
	}
	if k > n.Key {
		return get(n.Right, k)
	}

	return n.Value, true
}

// size returns the amount of nodes in the subtree rooted at n
func size[K cmp.Ordered, V any](n *Node[K, V]) int {
	if n == nil {
		return 0
	}

	return size(n.Left) + 1 + size(n.Right)
}

// getWithinRange returns all V values whose keys are between l and r (inclusive).
func getWithinRange[K cmp.Ordered, V any](n *Node[K, V], l, r K) (ret []V) {
	if n == nil {
		return ret
	}

	// If current node's key is greater than left bound, left subtree may have valid keys
	if n.Key > l {
		ret = append(ret, getWithinRange(n.Left, l, r)...)
	}

	// If current node's key is within range, include its value
	if n.Key >= l && n.Key <= r {
		ret = append(ret, n.Value)
	}

	// If current node's key is less than right bound, right subtree may have valid keys
	if n.Key < r {
		ret = append(ret, getWithinRange(n.Right, l, r)...)
	}

	return ret
}

// min returns the node with the smallest key in the subtree rooted at n
func max[K cmp.Ordered, V any](n *Node[K, V]) *Node[K, V] {
	if n.Right == nil {
		return n
	}

	return max(n.Right)
}

// min returns the node with the smallest key in the subtree rooted at n
func min[K cmp.Ordered, V any](n *Node[K, V]) *Node[K, V] {
	if n.Left == nil {
		return n
	}

	return min(n.Left)
}

func stringRepr[K cmp.Ordered, V any](n *Node[K, V], depth int) string {
	if n == nil {
		return ""
	}

	rTree := stringRepr(n.Right, depth+1)
	curr := fmt.Sprintf("%s[%v:%v]\n", strings.Repeat("\t", depth), n.Key, n.Value)
	lTree := stringRepr(n.Left, depth+1)

	return fmt.Sprintf("%s%s%s", rTree, curr, lTree)
}
