package tree

import (
	"reflect"
	"testing"
)

func isBalanced(node *AVLNode) bool {
	if node == nil {
		return true
	}
	balance := getBalanceFactor(node)
	if balance < -1 || balance > 1 {
		return false
	}
	return isBalanced(node.Left) && isBalanced(node.Right)
}

func TestAVLTree_InsertInOrder(t *testing.T) {
	tree := &AVLTree{}
	vals := []int{1, 2, 3, 4, 5, 6, 7}

	for _, v := range vals {
		tree.Insert(v)
	}

	got := tree.InOrderTraversal()
	if !reflect.DeepEqual(got, vals) {
		t.Errorf("InOrderTraversal mismatch: got %v, want %v", got, vals)
	}

	if !isBalanced(tree.Root) {
		t.Errorf("Tree is not balanced after in-order insertions")
	}
}

func TestAVLTree_InsertReverseOrder(t *testing.T) {
	tree := &AVLTree{}
	vals := []int{7, 6, 5, 4, 3, 2, 1}

	for _, v := range vals {
		tree.Insert(v)
	}

	want := []int{1, 2, 3, 4, 5, 6, 7}
	got := tree.InOrderTraversal()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("InOrderTraversal mismatch: got %v, want %v", got, want)
	}

	if !isBalanced(tree.Root) {
		t.Errorf("Tree is not balanced after reverse-order insertions")
	}
}

func TestAVLTree_InsertWithDuplicates(t *testing.T) {
	tree := &AVLTree{}
	vals := []int{10, 20, 30, 20, 10, 40, 50}

	for _, v := range vals {
		tree.Insert(v)
	}

	got := tree.InOrderTraversal()
	want := []int{10, 20, 30, 40, 50}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Duplicate handling failed: got %v, want %v", got, want)
	}
}

func TestAVLTree_InsertCausesLR(t *testing.T) {
	tree := &AVLTree{}
	// Insert values that cause Left-Right (LR) rotation
	vals := []int{30, 10, 20}

	for _, v := range vals {
		tree.Insert(v)
	}

	got := tree.InOrderTraversal()
	want := []int{10, 20, 30}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("LR rotation test failed: got %v, want %v", got, want)
	}

	if !isBalanced(tree.Root) {
		t.Errorf("Tree is not balanced after LR rotation case")
	}
}

func TestAVLTree_InsertCausesRL(t *testing.T) {
	tree := &AVLTree{}
	// Insert values that cause Right-Left (RL) rotation
	vals := []int{10, 30, 20}

	for _, v := range vals {
		tree.Insert(v)
	}

	got := tree.InOrderTraversal()
	want := []int{10, 20, 30}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("RL rotation test failed: got %v, want %v", got, want)
	}

	if !isBalanced(tree.Root) {
		t.Errorf("Tree is not balanced after RL rotation case")
	}
}
