package leetcode

func searchRotate(nums []int, target int) int {
	for i := range nums {
		if nums[i] == target {
			return i
		}
	}

	return -1
}
