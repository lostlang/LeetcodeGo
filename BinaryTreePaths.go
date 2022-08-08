package leetcode

import "fmt"

func binaryTreePaths(root *TreeNode) []string {
	return binaryTreePathsWithString(root, "")
}

func binaryTreePathsWithString(root *TreeNode, str string) []string {
	if root == nil {
		return []string{}
	}

	if root.Left == nil && root.Right == nil {
		return []string{str + fmt.Sprint(root.Val)}
	}

	left := binaryTreePathsWithString(root.Left, str+fmt.Sprint(root.Val)+"->")
	right := binaryTreePathsWithString(root.Right, str+fmt.Sprint(root.Val)+"->")

	return append(left, right...)
}
