type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func diameterOfBinaryTree(root *TreeNode) int {
  _, max := binaryTree(root)
  return max
}

func binaryTree(root *TreeNode) (int, int) {
	if root == nil {
		return 0,0;
	}
	leftH, leftMax := binaryTree(root.Left)
	rightH, rightMax := binaryTree(root.Right)
	return Max(leftH+1, rightH+1), Max(Max(leftMax,rightMax),leftH+rightH)
}