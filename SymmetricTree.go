package leetcode

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var curLeft = []*TreeNode{root.Left}
	var curRight = []*TreeNode{root.Right}

	for len(curLeft) != 0 {
		for i := range curLeft {
			if curLeft[i] == curRight[len(curRight)-i-1] {
				continue
			}

			if curLeft[i] == nil || curRight[len(curRight)-i-1] == nil {
				return false
			}

			if curLeft[i].Val != curRight[len(curRight)-i-1].Val {
				return false
			}
		}

		curLeft = getChildTreeNodes(curLeft)
		curRight = getChildTreeNodes(curRight)
	}

	return true
}

func getChildTreeNodes(nodes []*TreeNode) (out []*TreeNode) {
	for _, node := range nodes {
		if node != nil {
			out = append(out, node.Left)
			out = append(out, node.Right)
		}
	}

	return
}
