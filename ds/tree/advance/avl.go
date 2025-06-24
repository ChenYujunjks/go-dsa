package advance

import "fmt"

// 定义本地 AVL 树节点结构（本地类型）
type AVLNode struct {
	Val    int
	Left   *AVLNode
	Right  *AVLNode
	Height int
}

// 获取高度（nil 节点高度为 0）
func height(n *AVLNode) int {
	if n == nil {
		return 0
	}
	return n.Height
}

// 更新节点高度
func updateHeight(n *AVLNode) {
	if n == nil {
		return
	}
	lh, rh := height(n.Left), height(n.Right)
	if lh > rh {
		n.Height = lh + 1
	} else {
		n.Height = rh + 1
	}
}

// 平衡因子 = 左高 - 右高
func (n *AVLNode) balanceFactor() int {
	return height(n.Left) - height(n.Right)
}

// 示例插入函数（未实现旋转）
func (n *AVLNode) InsertAVL(val int) *AVLNode {
	if n == nil {
		return &AVLNode{Val: val, Height: 1}
	}

	if val < n.Val {
		n.Left = n.Left.InsertAVL(val)
	} else if val > n.Val {
		n.Right = n.Right.InsertAVL(val)
	} else {
		return n // 忽略重复插入
	}

	updateHeight(n)
	// 后续可在此处调用 balance & rotate
	return n
}

// 演示入口
func ExecuteAVL() {
	var root *AVLNode
	values := []int{10, 5, 15, 3, 7}
	for _, v := range values {
		root = root.InsertAVL(v)
	}
	fmt.Println("AVL Tree built (rotation logic not yet implemented)")
}
