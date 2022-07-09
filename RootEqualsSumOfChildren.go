package leetcode

func checkTree(root *TreeNode) bool {
	return root.Val == root.Left.Val+root.Right.Val
}
