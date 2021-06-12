/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 // 先翻转左子树，再判断左右子树是否相同
 func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
		return nil
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left = right
	root.Right = left
	return root
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	} else {
		return false
	}
}

func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }
    left := invertTree(root.Left)
    
    return isSameTree(left, root.Right)
}

// 解法一 dfs
func isSymmetric(root *TreeNode) bool {
	return root == nil || dfs(root.Left, root.Right)
}

func dfs(rootLeft, rootRight *TreeNode) bool {
	if rootLeft == nil && rootRight == nil {
		return true
	}
	if rootLeft == nil || rootRight == nil {
		return false
	}
	if rootLeft.Val != rootRight.Val {
		return false
	}
	return dfs(rootLeft.Left, rootRight.Right) && dfs(rootLeft.Right, rootRight.Left)
}