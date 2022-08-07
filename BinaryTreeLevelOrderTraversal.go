package leetcode

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var out [][]int
	var level []*TreeNode
	var nextLevel = []*TreeNode{root}

	for len(nextLevel) > 0 {
		level = nextLevel
		nextLevel = []*TreeNode{}
		out = append(out, []int{})

		for _, node := range level {
			out[len(out)-1] = append(out[len(out)-1], node.Val)

			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
	}

	return out
}
