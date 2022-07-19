package leetcode

func search(nums []int, target int) int {
	var left = 0
	var right = len(nums) - 1

	for left <= right {
		var pivot = left + (right-left)/2

		if nums[pivot] > target {
			right = pivot - 1
		} else if nums[pivot] < target {
			left = pivot + 1
		} else {
			return pivot
		}
	}

	return -1
}
