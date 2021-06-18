/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCousins(root *TreeNode, x int, y int) bool {
	var depth1, depth2, parent1, parent2 int
	dfsCousins(root, x, 0, -1, &parent1, &depth1)
	dfsCousins(root, y, 0, -1, &parent2, &depth2)
	return depth1 > 1 && depth1 == depth2 && parent1 != parent2
}

func dfsCousins(root *TreeNode, val, depth, last int, parent, res *int) {
	if root == nil {
		return
	}
	if root.Val == val {
		*res = depth
		*parent = last
		return
	}
	depth++
	dfsCousins(root.Left, val, depth, root.Val, parent, res)
	dfsCousins(root.Right, val, depth, root.Val, parent, res)
}