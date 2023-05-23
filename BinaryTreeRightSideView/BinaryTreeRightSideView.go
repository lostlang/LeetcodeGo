package BinaryTreeRightSideView

import "leetcode/utils"

type TreeNode = utils.TreeNode

func rightSideView(root *TreeNode) []int {
	out := []int{}

	if root == nil {
		return out
	}

	currentLevel := []*TreeNode{root}
	nextLevel := []*TreeNode{}

	for len(currentLevel) != 0 {

		for _, node := range currentLevel {
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}

		out = append(out, currentLevel[len(currentLevel)-1].Val)

		currentLevel = nextLevel
		nextLevel = []*TreeNode{}

	}

	return out
}
