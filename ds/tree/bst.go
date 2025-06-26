package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (node *TreeNode) isLeaf() bool {
	if node == nil {
		return false // 或者 panic("nil node")，看你需要
	}
	return node.Left == nil && node.Right == nil
}

type BST struct {
	Root *TreeNode
}

// 严格不重复
func (t *BST) IsValid() bool {
	return validate(t.Root, nil, nil)
}

func validate(node *TreeNode, min *int, max *int) bool {
	if node == nil {
		return true
	}
	if min != nil && node.Val <= *min {
		return false
	}
	if max != nil && node.Val >= *max {
		return false
	}
	return validate(node.Left, min, &node.Val) && validate(node.Right, &node.Val, max)
}

func (t *BST) Search(val int) *TreeNode { return searchNode(t.Root, val) }
func searchNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if root.Val < val {
		return searchNode(root.Right, val)
	}
	return searchNode(root.Left, val)
}

func (t *BST) Insert(val int) {
	t.Root = insertRecursive(t.Root, val)
}

func insertRecursive(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return &TreeNode{Val: val}
	}

	if val < node.Val {
		node.Left = insertRecursive(node.Left, val)
	} else if val > node.Val {
		node.Right = insertRecursive(node.Right, val)
	}
	// 若 val == node.Val，忽略重复插入（也可以选择允许重复）
	return node
}

// Delete 删除节点
func (t *BST) Delete(key int) {
	t.Root = deleteNode(t.Root, key)
}
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else {
		// ✅ 找到要删除的节点（val == key）

		// 情况 1：叶子节点（左右都是 nil）
		if root.isLeaf() {
			return nil
		}

		// 情况 2：只有一个子节点
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}

		// 定义 辅助函数
		findMin := func(node *TreeNode) *TreeNode {
			for node.Left != nil {
				node = node.Left
			}
			return node
		}
		minNode := findMin(root.Right)
		root.Val = minNode.Val
		root.Right = deleteNode(root.Right, minNode.Val)
	}
	return root
}

// getTreeDepth calculates the depth of the tree
func (t *BST) Height() int {
	return MaxDepth(t.Root)
}
