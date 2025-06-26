package avl

// InOrderTraversal 中序遍历（升序）
func (t *AVLTree) InOrderTraversal() []int {
	var result []int
	inOrder(t.Root, &result)
	return result
}

// avl inOrder
func inOrder(n *AVLNode, result *[]int) {
	if n == nil {
		return
	}
	inOrder(n.Left, result)
	*result = append(*result, n.Val)
	inOrder(n.Right, result)
}

// rotateLeft 左旋
func rotateLeft(z *AVLNode) *AVLNode {
	y := z.Right
	T2 := y.Left

	y.Left = z
	z.Right = T2

	z.Height = 1 + max(height(z.Left), height(z.Right))
	y.Height = 1 + max(height(y.Left), height(y.Right))

	return y
}

// rotateRight 右旋
func rotateRight(y *AVLNode) *AVLNode {
	x := y.Left
	T2 := x.Right

	x.Right = y
	y.Left = T2

	y.Height = 1 + max(height(y.Left), height(y.Right))
	x.Height = 1 + max(height(x.Left), height(x.Right))

	return x
}
