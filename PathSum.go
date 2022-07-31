package leetcode

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	targetSum -= root.Val

	if targetSum == 0 && root.Left == root.Right {
		return true
	}

	return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
}
