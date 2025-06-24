package tree

import (
	"testing"
)

// TestBSTOperations 测试 BST 的插入、查找、删除、合法性校验
func TestBSTOperations(t *testing.T) {
	// 创建一个新的 BST
	bst := &BST{}

	// 插入一组节点
	vals := []int{50, 30, 70, 20, 40, 60, 80}
	for _, val := range vals {
		bst.Insert(val)
	}

	// 验证是否构成合法 BST
	if !bst.IsValid() {
		t.Errorf("Expected BST to be valid after insertions")
	}

	// 测试查找存在的节点
	node := bst.Search(60)
	if node == nil || node.Val != 60 {
		t.Errorf("Expected to find node with value 60")
	}

	// 测试查找不存在的节点
	node = bst.Search(100)
	if node != nil {
		t.Errorf("Expected not to find node with value 100")
	}

	// 删除一个叶子节点（20）
	bst.Delete(20)
	if bst.Search(20) != nil {
		t.Errorf("Expected node 20 to be deleted")
	}

	// 删除一个只有一个子节点的节点（30）
	bst.Delete(30)
	if bst.Search(30) != nil {
		t.Errorf("Expected node 30 to be deleted")
	}

	// 删除一个有两个子节点的节点（50）
	bst.Delete(50)
	if bst.Search(50) != nil {
		t.Errorf("Expected node 50 to be deleted")
	}

	// 删除后再次校验合法性
	if !bst.IsValid() {
		t.Errorf("Expected BST to be valid after deletions")
	}
}
func TestBSTHeight(t *testing.T) {
	bst := &BST{}
	vals := []int{10, 5, 1, 7, 15, 12, 18}
	for _, val := range vals {
		bst.Insert(val)
	}

	expectedHeight := 3 // 完全平衡树高度 = 3
	if h := bst.Height(); h != expectedHeight {
		t.Errorf("Expected height %d, but got %d", expectedHeight, h)
	}

	bst = &BST{}
	if h := bst.Height(); h != 0 {
		t.Errorf("Expected height 0 for empty tree, but got %d", h)
	}
}
