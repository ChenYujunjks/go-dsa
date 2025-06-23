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
func IsValidBST(root *TreeNode) bool {
	return validate(root, nil, nil)
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

func SearchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if root.Val < val {
		return SearchBST(root.Right, val)
	}
	return SearchBST(root.Left, val)
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

func DeleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if key < root.Val {
		root.Left = DeleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = DeleteNode(root.Right, key)
	} else {
		// ✅ 找到要删除的节点（val == key）

		// 情况 1：叶子节点（左右都是 nil）
		if root.Left == nil && root.Right == nil {
			return nil
		}

		// 情况 2：只有一个子节点
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}

		// 情况 3 留空，后续处理
	}
}