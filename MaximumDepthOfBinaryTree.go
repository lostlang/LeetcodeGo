package leetcode

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var leftRoot = maxDepth(root.Left)
	var rightRoot = maxDepth(root.Right)

	if leftRoot > rightRoot {
		return leftRoot + 1
	}

	return rightRoot + 1
}
