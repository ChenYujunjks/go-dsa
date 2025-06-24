package advance

import "fmt"

// 定义颜色类型
const (
	Red   = true
	Black = false
)

type RBNode struct {
	Val    int
	Color  bool // true 表示红色，false 表示黑色
	Left   *RBNode
	Right  *RBNode
	Parent *RBNode
}

type RBTree struct {
	Root *RBNode
}

// NewRBNode 创建一个红色的新节点
func NewRBNode(val int) *RBNode {
	return &RBNode{
		Val:   val,
		Color: Red,
	}
}

// Insert 插入新值（暂不包含修复逻辑）
func (t *RBTree) Insert(val int) {
	n := NewRBNode(val)
	if t.Root == nil {
		n.Color = Black // 根节点必须是黑色
		t.Root = n
		return
	}
	insertBST(t.Root, n)
	// TODO: 插入修复操作（红黑树平衡）
}

// insertBST 按照普通 BST 的方式插入新节点（无旋转）
func insertBST(root, node *RBNode) {
	if node.Val < root.Val {
		if root.Left == nil {
			root.Left = node
			node.Parent = root
		} else {
			insertBST(root.Left, node)
		}
	} else {
		if root.Right == nil {
			root.Right = node
			node.Parent = root
		} else {
			insertBST(root.Right, node)
		}
	}
}

// InOrderTraversal 中序遍历调试函数
func (t *RBTree) InOrderTraversal() {
	var dfs func(node *RBNode)
	dfs = func(node *RBNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		fmt.Printf("%d(%s) ", node.Val, colorStr(node.Color))
		dfs(node.Right)
	}
	dfs(t.Root)
	fmt.Println()
}

func colorStr(c bool) string {
	if c {
		return "R"
	}
	return "B"
}

// Demo 执行入口
func DemoRB() {
	t := &RBTree{}
	for _, val := range []int{10, 5, 15, 3, 7, 12, 18} {
		t.Insert(val)
	}
	t.InOrderTraversal()
}
