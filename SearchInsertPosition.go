package leetcode

func searchInsert(nums []int, target int) int {
	var left int
	var right = len(nums) - 1
	var pivot int

	for left <= right {
		pivot = left + (right-left)/2

		if nums[pivot] > target {
			right = pivot - 1
		} else if nums[pivot] < target {
			left = pivot + 1
		} else {
			return pivot
		}
	}

	if nums[pivot] > target {
		return pivot
	} else {
		return pivot + 1
	}

}
