package leetcode

func reverseOddLevels(root *TreeNode) *TreeNode {
	var level int
	var current = []*TreeNode{root}
	var next []*TreeNode

	for len(current) > 0 {
		level++

		for _, node := range current {
			if node.Left != nil {
				next = append(next, node.Left)
			}

			if node.Right != nil {
				next = append(next, node.Right)
			}
		}

		if level%2 == 1 {
			for i := 0; i < len(next)/2; i++ {
				next[i].Val, next[len(next)-i-1].Val = next[len(next)-i-1].Val, next[i].Val
			}
		}

		current = next
		next = nil
	}

	return root
}
