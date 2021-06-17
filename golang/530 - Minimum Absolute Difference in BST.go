/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {
	res := 100000
	pre := -1
	dfs(root, &res, &pre)
	return res
}

func dfs(root *TreeNode, res, pre *int) {
	if root == nil {
		return
	}

	dfs(root.Left, res, pre)
	if *pre != -1 {
		*res = Min(*res, Abs(root.Val-*pre))
	}
	*pre = root.Val
	dfs(root.Right, res, pre)
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}