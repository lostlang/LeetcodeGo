package leetcode

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max3(maxDepth(root.Left)+maxDepth(root.Right), diameterOfBinaryTree(root.Left), diameterOfBinaryTree(root.Right))
}

func max3(a, b, c int) int {
	if a > b {
		if a > c {
			return a
		}
		return c
	}
	if b > c {
		return b
	}
	return c
}
