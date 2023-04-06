package MaximumDepthOfBinaryTree

import "leetcode/utils"

type TreeNode = utils.TreeNode

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftRoot := maxDepth(root.Left)
	rightRoot := maxDepth(root.Right)

	if leftRoot > rightRoot {
		return leftRoot + 1
	}

	return rightRoot + 1
}

var MaxDepth = maxDepth
