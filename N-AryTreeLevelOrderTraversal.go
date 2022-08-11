package leetcode

func levelOrderNodes(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}

	var out [][]int
	var nextLevel []*Node
	var curLevel = []*Node{root}

	for len(curLevel) > 0 {
		nextLevel = []*Node{}
		out = append(out, []int{})

		for _, node := range curLevel {
			if node != nil {
				out[len(out)-1] = append(out[len(out)-1], node.Val)
			}
			nextLevel = append(nextLevel, node.Children...)
		}

		curLevel = nextLevel
	}

	return out
}
