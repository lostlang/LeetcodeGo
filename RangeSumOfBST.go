package leetcode

func rangeSumBST(root *TreeNode, low int, high int) int {
	output := 0
	if root == nil {
		return output
	}
	if root.Left != nil {
		output += rangeSumBST(root.Left, low, high)
	}
	if root.Right != nil {
		output += rangeSumBST(root.Right, low, high)
	}
	if root.Val >= low && root.Val <= high {
		output += root.Val
	}

	return output
}
