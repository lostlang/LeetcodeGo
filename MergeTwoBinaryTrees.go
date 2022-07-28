package leetcode

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	var out = &TreeNode{0, nil, nil}

	if root1 == nil {
		return root2
	}

	if root2 == nil {
		return root1
	}

	out.Val = root1.Val + root2.Val
	out.Left = mergeTrees(root1.Left, root2.Left)
	out.Right = mergeTrees(root1.Right, root2.Right)

	return out
}
