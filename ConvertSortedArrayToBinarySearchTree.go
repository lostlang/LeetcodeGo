package leetcode

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 3 {
		return &TreeNode{nums[1], &TreeNode{nums[0], nil, nil}, &TreeNode{nums[2], nil, nil}}
	} else if len(nums) == 2 {
		return &TreeNode{nums[1], &TreeNode{nums[0], nil, nil}, nil}
	} else if len(nums) == 1 {
		return &TreeNode{nums[0], nil, nil}
	}

	center := len(nums) / 2

	return &TreeNode{nums[center], sortedArrayToBST(nums[:center]), sortedArrayToBST(nums[center+1:])}
}
