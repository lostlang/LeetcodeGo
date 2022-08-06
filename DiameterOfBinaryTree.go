package leetcode

func diameterOfBinaryTree(root *TreeNode) int {
	return maxDepth(root.Left) + maxDepth(root.Right)
}
