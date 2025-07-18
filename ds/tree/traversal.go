package tree

import "github.com/ChenYujunjks/go-review/ds/queue"

// 辅助递归函数，使用指针收集结果
func inorder(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}
	inorder(node.Left, result)
	*result = append(*result, node.Val)
	inorder(node.Right, result)
}
func preorder(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}
	*result = append(*result, node.Val)
	preorder(node.Left, result)
	preorder(node.Right, result)
}
func postorder(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}
	postorder(node.Left, result)
	postorder(node.Right, result)
	*result = append(*result, node.Val)
}

// InorderTraversal 返回 BST 的中序遍历结果：左 -> 根 -> 右
func (t *BST) InorderTraversal() []int {
	var result []int
	inorder(t.Root, &result)
	return result
}

// PreorderTraversal 返回 BST 的前序遍历结果：根 -> 左 -> 右
func (t *BST) PreorderTraversal() []int {
	var result []int
	preorder(t.Root, &result)
	return result
}

func (t *BST) PostorderTraversal() []int {
	var result []int
	postorder(t.Root, &result)
	return result
}

// bfs
func (root *TreeNode) LevelOrderTraversal() []int {
	var result []int
	if root == nil {
		return result
	}
	q := queue.New()
	q.Enqueue(root)

	for !q.IsEmpty() {
		nodeAny, _ := q.Dequeue()
		node := nodeAny.(*TreeNode) //类型断言（type assertion）回原始类型
		//“我确信接口变量 x 中存的是 T 类型的值，请把它强制还原成 T 类型。”
		result = append(result, node.Val)
		if node.Left != nil {
			q.Enqueue(node.Left)
		}
		if node.Right != nil {
			q.Enqueue(node.Right)
		}
	}
	return result
}
