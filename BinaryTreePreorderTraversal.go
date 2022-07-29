package leetcode

func preorderTraversal(root *TreeNode) []int {
	var out = make([]int, 0)

	if root == nil {
		return out
	}

	out = append(out, root.Val)

	out = append(out, preorderTraversal(root.Left)...)
	out = append(out, preorderTraversal(root.Right)...)

	return out
}
