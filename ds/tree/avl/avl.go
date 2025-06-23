package avl

import (
	"fmt"

	"github.com/ChenYujunjks/go-review/ds/tree"
	"github.com/ChenYujunjks/go-review/ds/tree/bst"
)

// Utility to get node height
func height(n *tree.TreeNode) int {
	if n == nil {
		return 0
	}
	return n.Height
}

// Update node height
func (n *tree.TreeNode) updateHeight() {
	lh, rh := height(n.Left), height(n.Right)
	if lh > rh {
		n.Height = lh + 1
	} else {
		n.Height = rh + 1
	}
}

// Calculate balance factor
func (n *tree.TreeNode) balanceFactor() int {
	return height(n.Left) - height(n.Right)
}

// Placeholder for rotation and insertion logic
func (n *tree.TreeNode) Insert_AVL(value int) *bst.Node {
	// AVL Insertion logic should be added here, including rotations
	// For now, this is a placeholder for demonstration.
	return n
}

// ExecuteAVL runs the AVL tree logic as a demo
func ExecuteAVL() {
	fmt.Println("AVL: Demo execution (rotation logic not implemented)")
}
