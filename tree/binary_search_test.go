package tree_test

import (
	"fmt"
	"testing"

	"github.com/zmolz/gofuncyourself/tree"
)

func TestInsert(t *testing.T) {
	t.Run("test_insert_root", func(t *testing.T) {
		tree := &tree.BST[int, string]{}
		tree.Insert(1, "root")

		if tree.Size() != 1 {
			t.Fatalf("Tree should only contain 1 element, got %d", tree.Size())
		}

		if tree.Root == nil {
			t.Fatal("Root should not be nil")
		}

		if tree.Root.Key != 1 || tree.Root.Value != "root" {
			t.Fatalf("Insert() failed: got root %+v", tree.Root)
		}
	})

	t.Run("test_insert_multilevel", func(t *testing.T) {
		tree := &tree.BST[int, string]{}
		tree.Insert(10, "root")
		tree.Insert(5, "left")
		tree.Insert(15, "right")
		tree.Insert(3, "left.left")
		tree.Insert(7, "left.right")
		tree.Insert(12, "right.left")
		tree.Insert(18, "right.right")

		if tree.Size() != 7 {
			t.Fatalf("Tree should contain 7 elements, got %d", tree.Size())
		}

		if tree.Root.Key != 10 {
			t.Fatalf("Root key should be 10, got %d", tree.Root.Key)
		}
		if tree.Root.Left == nil || tree.Root.Left.Key != 5 {
			t.Fatalf("Left child of root should be 5, got %+v", tree.Root.Left)
		}
		if tree.Root.Right == nil || tree.Root.Right.Key != 15 {
			t.Fatalf("Right child of root should be 15, got %+v", tree.Root.Right)
		}
	})
}

func TestInsert_Update(t *testing.T) {
	t.Run("insert_existing_key_updates_value", func(t *testing.T) {
		tree := &tree.BST[int, string]{}
		tree.Insert(10, "root")

		// Insert same key with different value
		tree.Insert(10, "new_root")

		if tree.Size() != 1 {
			t.Fatalf("Tree should still contain 1 element, got %d", tree.Size())
		}

		val, ok := tree.Get(10)
		if !ok {
			t.Fatal("Expected key 10 to exist")
		}

		if val != "new_root" {
			t.Fatalf("Expected value 'new_root' for key 10, got %q", val)
		}
	})
}

func TestGet(t *testing.T) {
	t.Run("test_get", func(t *testing.T) {
		tree := &tree.BST[int, string]{}
		tree.Insert(10, "root")
		tree.Insert(5, "left")
		tree.Insert(15, "right")
		tree.Insert(3, "left.left")
		tree.Insert(7, "left.right")
		tree.Insert(12, "right.left")
		tree.Insert(18, "right.right")

		tests := []struct {
			key      int
			expected string
			found    bool
		}{
			{10, "root", true},
			{5, "left", true},
			{15, "right", true},
			{3, "left.left", true},
			{7, "left.right", true},
			{12, "right.left", true},
			{18, "right.right", true},
			{99, "", false}, // non-existent key
		}

		for _, tt := range tests {
			val, ok := tree.Get(tt.key)
			if ok != tt.found {
				t.Fatalf("Get(%d): expected found=%v, got %v", tt.key, tt.found, ok)
			}
			if ok && val != tt.expected {
				t.Fatalf("Get(%d): expected value=%q, got %q", tt.key, tt.expected, val)
			}
		}
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete_leaf_node", func(t *testing.T) {
		tree := &tree.BST[int, string]{}
		tree.Insert(10, "root")
		tree.Insert(5, "left")
		tree.Insert(15, "right")

		deleted := tree.Delete(5)
		if !deleted {
			t.Fatal("Expected Delete(5) to return true")
		}

		if _, ok := tree.Get(5); ok {
			t.Fatal("Expected key 5 to be removed")
		}
		if tree.Size() != 2 {
			t.Fatalf("Expected size=2 after delete, got %d", tree.Size())
		}
	})

	t.Run("delete_node_with_one_child", func(t *testing.T) {
		tree := &tree.BST[int, string]{}
		tree.Insert(10, "root")
		tree.Insert(5, "left")
		tree.Insert(3, "left.left")

		deleted := tree.Delete(5)
		if !deleted {
			t.Fatal("Expected Delete(5) to return true")
		}
		if _, found := tree.Get(5); found {
			t.Fatal("Expected key 5 to be removed")
		}
		if _, found := tree.Get(3); !found {
			t.Fatal("Expected key 3 to remain after delete")
		}
		if tree.Size() != 2 {
			t.Fatalf("Expected size=2 after delete, got %d", tree.Size())
		}
	})

	t.Run("delete_node_with_two_children", func(t *testing.T) {
		tree := &tree.BST[int, string]{}
		tree.Insert(10, "root")
		tree.Insert(5, "left")
		tree.Insert(15, "right")
		tree.Insert(12, "right.left")
		tree.Insert(18, "right.right")

		fmt.Printf("%s\n", tree)
		deleted := tree.Delete(15)
		fmt.Printf("%s\n", tree)
		if !deleted {
			t.Fatal("Expected Delete(15) to return true")
		}
		if _, found := tree.Get(15); found {
			t.Fatal("Expected key 15 to be removed")
		}
		if tree.Size() != 4 {
			t.Fatalf("Expected size=4 after delete, got %d", tree.Size())
		}
	})

	t.Run("delete_nonexistent", func(t *testing.T) {
		tree := &tree.BST[int, string]{}
		tree.Insert(10, "root")

		deleted := tree.Delete(99)
		if deleted {
			t.Fatal("Expected Delete(99) to return false for non-existent key")
		}
	})
}

func TestMinMax(t *testing.T) {
	tree := &tree.BST[int, string]{}
	tree.Insert(10, "root")
	tree.Insert(5, "left")
	tree.Insert(15, "right")
	tree.Insert(3, "min")
	tree.Insert(20, "max")

	minKey, minVal, ok := tree.Min()
	if !ok || minKey != 3 || minVal != "min" {
		t.Fatalf("Min() failed: expected (3, 'min'), got (%v, %q, %v)", minKey, minVal, ok)
	}

	maxKey, maxVal, ok := tree.Max()
	if !ok || maxKey != 20 || maxVal != "max" {
		t.Fatalf("Max() failed: expected (20, 'max'), got (%v, %q, %v)", maxKey, maxVal, ok)
	}
}

func TestSize(t *testing.T) {
	tree := &tree.BST[int, string]{}
	if tree.Size() != 0 {
		t.Fatalf("Expected empty tree size=0, got %d", tree.Size())
	}

	tree.Insert(10, "root")
	tree.Insert(5, "left")
	tree.Insert(15, "right")

	if tree.Size() != 3 {
		t.Fatalf("Expected size=3, got %d", tree.Size())
	}

	tree.Delete(15)
	if tree.Size() != 2 {
		t.Fatalf("Expected size=2 after delete, got %d", tree.Size())
	}
}
