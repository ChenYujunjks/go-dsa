package tree

import (
	"reflect"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	bst := &BST{}
	values := []int{5, 3, 7, 2, 4, 6, 8}
	for _, v := range values {
		bst.Insert(v)
	}

	expected := []int{2, 3, 4, 5, 6, 7, 8}
	actual := bst.InorderTraversal()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestPreorderTraversal(t *testing.T) {
	bst := &BST{}
	values := []int{5, 3, 7, 2, 4, 6, 8}
	for _, v := range values {
		bst.Insert(v)
	}

	expected := []int{5, 3, 2, 4, 7, 6, 8}
	actual := bst.PreorderTraversal()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestPostorderTraversal(t *testing.T) {
	bst := &BST{}
	values := []int{5, 3, 7, 2, 4, 6, 8}
	for _, v := range values {
		bst.Insert(v)
	}

	expected := []int{2, 4, 3, 6, 8, 7, 5}
	actual := bst.PostorderTraversal()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
func TestLevelOrderTraversal(t *testing.T) {
	bst := &BST{}
	values := []int{5, 3, 7, 2, 4, 8}
	for _, v := range values {
		bst.Insert(v)
	}

	expected := []int{5, 3, 7, 2, 4, 8}
	actual := bst.Root.LevelOrderTraversal()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestInvertTree(t *testing.T) {
	// 原始树：
	//     4
	//    / \
	//   2   7
	//  / \ / \
	// 1  3 6  9
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{
			Val:   7,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 9},
		},
	}

	inverted := InvertTree(root)

	expected := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   7,
			Left:  &TreeNode{Val: 9},
			Right: &TreeNode{Val: 6},
		},
		Right: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 1},
		},
	}

	if !reflect.DeepEqual(inverted, expected) {
		t.Errorf("InvertTree failed. Expected %v, got %v", expected, inverted)
	}
}
func TestIsSymmetric(t *testing.T) {
	// 对称的树：
	//     1
	//    / \
	//   2   2
	//  / \ / \
	// 3  4 4  3
	symmetricRoot := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 3},
		},
	}

	if !IsSymmetric(symmetricRoot) {
		t.Errorf("IsSymmetric failed on symmetric tree")
	}

	// 非对称的树：
	//     1
	//    / \
	//   2   2
	//    \   \
	//     3   3
	asymmetricRoot := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{
			Val:   2,
			Right: &TreeNode{Val: 3},
		},
	}

	if IsSymmetric(asymmetricRoot) {
		t.Errorf("IsSymmetric incorrectly returned true for asymmetric tree")
	}
}
func LevelOrderTraversalFlat(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		result = append(result, node.Val)
		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}
	}
	return result
}

func TestBuildTree_NormalCase(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}

	root := BuildTree(preorder, inorder)

	got := LevelOrderTraversal(root) // ✅ 使用你自己的 BFS 函数
	want := []int{3, 9, 20, 15, 7}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("BuildTree failed. Expected %v, got %v", want, got)
	}
}

func TestBuildTree_LeftOnly(t *testing.T) {
	preorder := []int{5, 4, 3, 2, 1}
	inorder := []int{1, 2, 3, 4, 5}

	root := BuildTree(preorder, inorder)

	got := LevelOrderTraversalFlat(root)
	want := []int{5, 4, 3, 2, 1}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("BuildTree (left only) failed. Expected %v, got %v", want, got)
	}
}
