package leetcode

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var left = maxDepth(root.Left)
	var right = maxDepth(root.Right)

	if left-right > -2 && left-right < 2 {
		return isBalanced(root.Left) && isBalanced(root.Right)
	}

	return false
}
