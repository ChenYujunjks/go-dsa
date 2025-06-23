package tree

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
	root.Left = BuildTree(preorder[1+idx:], inorder[idx+1:])
	return root
}
