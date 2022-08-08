package leetcode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pInLeft := searchTreeNode(root.Left, p)
	qInLeft := searchTreeNode(root.Left, q)

	pInRight := searchTreeNode(root.Right, p)
	qInRight := searchTreeNode(root.Right, q)

	if pInLeft && qInLeft {
		return lowestCommonAncestor(root.Left, p, q)
	}

	if pInRight && qInRight {
		return lowestCommonAncestor(root.Right, p, q)
	}

	return root
}

func searchTreeNode(root, node *TreeNode) bool {
	if root == nil {
		return false
	}

	if root == node {
		return true
	}

	return searchTreeNode(root.Left, node) || searchTreeNode(root.Right, node)
}
