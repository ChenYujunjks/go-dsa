package tree

import (
	"reflect"
	"testing"
)

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

func TestBuildTree_NormalCase(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}

	root := BuildTree(preorder, inorder)

	got := root.LevelOrderTraversal() // ✅ 使用你自己的 BFS 函数
	want := []int{3, 9, 20, 15, 7}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("BuildTree failed. Expected %v, got %v", want, got)
	}
}

func TestBuildTree_LeftOnly(t *testing.T) {
	preorder := []int{5, 4, 3, 2, 1}
	inorder := []int{1, 2, 3, 4, 5}

	root := BuildTree(preorder, inorder)

	got := root.LevelOrderTraversal()
	want := []int{5, 4, 3, 2, 1}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("BuildTree (left only) failed. Expected %v, got %v", want, got)
	}
}
func TestBuildTree_Empty(t *testing.T) {
	var preorder []int
	var inorder []int

	root := BuildTree(preorder, inorder)

	if root != nil {
		t.Errorf("BuildTree (empty) failed. Expected nil, got non-nil")
	}
}
