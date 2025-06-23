package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BST struct {
	Root *TreeNode
}

func (node *TreeNode) isLeaf() bool {
	if node == nil {
		return false // 或者 panic("nil node")，看你需要
	}
	return node.Left == nil && node.Right == nil
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
