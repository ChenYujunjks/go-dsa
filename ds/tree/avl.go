package tree

type AVLNode struct {
	Val    int
	Left   *AVLNode
	Right  *AVLNode
	Height int
}

type AVLTree struct {
	Root *AVLNode
}

// height 返回节点高度，空节点为 -1
func height(n *AVLNode) int {
	if n == nil {
		return -1
	}
	return n.Height
}

// getBalanceFactor 返回节点的平衡因子
func getBalanceFactor(n *AVLNode) int {
	if n == nil {
		return 0
	}
	return height(n.Left) - height(n.Right)
}

// Insert 向 AVL 树中插入一个值
func (t *AVLTree) Insert(val int) {
	t.Root = insert(t.Root, val)
}

func insert(node *AVLNode, val int) *AVLNode {
	if node == nil {
		return &AVLNode{Val: val, Height: 0}
	}

	if val < node.Val {
		node.Left = insert(node.Left, val)
	} else if val > node.Val {
		node.Right = insert(node.Right, val)
	} else {
		return node // 不允许重复值
	}

	// 更新高度
	node.Height = 1 + max(height(node.Left), height(node.Right))

	// 获取平衡因子并旋转
	balance := getBalanceFactor(node)

	// 四种旋转情况
	if balance > 1 && val < node.Left.Val {
		return rotateRight(node) // LL
	}
	if balance < -1 && val > node.Right.Val {
		return rotateLeft(node) // RR
	}
	if balance > 1 && val > node.Left.Val {
		node.Left = rotateLeft(node.Left) // LR
		return rotateRight(node)
	}
	if balance < -1 && val < node.Right.Val {
		node.Right = rotateRight(node.Right) // RL
		return rotateLeft(node)
	}

	return node
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

// InOrderTraversal 中序遍历（升序）
func (t *AVLTree) InOrderTraversal() []int {
	var result []int
	inOrder(t.Root, &result)
	return result
}

func inOrder(n *AVLNode, result *[]int) {
	if n == nil {
		return
	}
	inOrder(n.Left, result)
	*result = append(*result, n.Val)
	inOrder(n.Right, result)
}
