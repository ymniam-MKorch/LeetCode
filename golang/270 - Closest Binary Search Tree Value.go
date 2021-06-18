package main
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func closestValue(root *TreeNode, target float64) int {
   
}

func cv(root *TreeNode, target float64) int, float64 {
	if root == nil {
		return 
	}
	leftC, leftDiff := cv(root.Left, target)
}