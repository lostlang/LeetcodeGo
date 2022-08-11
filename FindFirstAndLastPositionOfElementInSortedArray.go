package leetcode

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	insertPosition := search(nums, target)

	if insertPosition == -1 {
		return []int{-1, -1}
	}

	left := insertPosition
	right := insertPosition

	for left > 1 && nums[left-1] == target {
		left--
	}

	for right < len(nums)-1 && nums[right+1] == target {
		right++
	}

	return []int{left, right}
}
