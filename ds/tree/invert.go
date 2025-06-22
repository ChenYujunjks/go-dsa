package tree

func InvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = InvertTree(root.Right), InvertTree(root.Left)
	return root
}

// IsSymmetric 判断二叉树是否左右对称
func IsSymmetric(root *TreeNode) bool {
	// 你来实现
}
