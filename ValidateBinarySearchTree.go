package leetcode

func isValidBST(root *TreeNode) bool {
	return isValidBSTWithMax(root, -1<<63, 1<<63-1)
}

func isValidBSTWithMax(root *TreeNode, left, right int) bool {
	if root == nil {
		return true
	}

	if root.Val <= left || root.Val >= right {
		return false
	}

	return isValidBSTWithMax(root.Left, left, root.Val) && isValidBSTWithMax(root.Right, root.Val, right)
}
