package BinaryTreeLevelOrderTraversal

import (
	"leetcode/utils"
)

type TreeNode = utils.TreeNode

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var output [][]int
	var level []*TreeNode
	nextLevel := []*TreeNode{root}

	for len(nextLevel) > 0 {
		level = nextLevel
		nextLevel = []*TreeNode{}
		output = append(output, []int{})

		for _, node := range level {
			output[len(output)-1] = append(output[len(output)-1], node.Val)

			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
	}

	return output
}
