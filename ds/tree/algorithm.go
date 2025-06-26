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
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}
func isMirror(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	if t1.Val != t2.Val {
		return false
	}
	return isMirror(t1.Left, t2.Right) && isMirror(t1.Right, t2.Left)
}

// BuildTree 根据 preorder 和 inorder 构建二叉树
func BuildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	rootVal := preorder[0]
	root := &TreeNode{Val: rootVal}
	// 在中序遍历中找到根节点的位置
	var idx int
	for i, v := range inorder {
		if v == rootVal {
			idx = i
			break
		}
	}
	//因为 左边的树的数量一定是相等的 所以 preorder[1:1+idx]和inorder[:idx] 都是左tree
	// preorder[0]和inorder[idx] 都是根节点
	root.Left = BuildTree(preorder[1:1+idx], inorder[:idx])
	root.Right = BuildTree(preorder[idx+1:], inorder[idx+1:])
	return root
}

func HasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	// 如果到了叶子节点
	if root.isLeaf() {
		return targetSum == root.Val
	}
	t := targetSum - root.Val

	return HasPathSum(root.Left, t) || HasPathSum(root.Right, t)
}

// MaxDepth 返回二叉树的最大深度
func MaxDepth(root *TreeNode) int {
	// 你来实现
	if root == nil {
		return 0
	}
	if root.isLeaf() {
		return 1
	}
	left := MaxDepth(root.Left)
	right := MaxDepth(root.Right)
	return 1 + max(left, right)
}

func balanceBST(root *TreeNode) *TreeNode {
	var nodes []*TreeNode

	// 中序遍历收集节点
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		nodes = append(nodes, node)
		inorder(node.Right)
	}
	inorder(root)

	// nodes已经按顺序写好 构建平衡 BST
	var build func(int, int) *TreeNode
	build = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}
		mid := (left + right) / 2
		root := nodes[mid]
		root.Left = build(left, mid-1)
		root.Right = build(mid+1, right)
		return root
	}

	return build(0, len(nodes)-1)
}
