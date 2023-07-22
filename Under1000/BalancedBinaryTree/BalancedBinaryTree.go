package BalancedBinaryTree

import (
	"leetcode/MaximumDepthOfBinaryTree"
	"leetcode/utils"
)

type TreeNode = utils.TreeNode

var maxDepth = MaximumDepthOfBinaryTree.MaxDepth

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	if left-right > -2 && left-right < 2 {
		return isBalanced(root.Left) && isBalanced(root.Right)
	}

	return false
}
