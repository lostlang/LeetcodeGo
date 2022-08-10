package leetcode

import (
	"sort"
)

func specialArray(nums []int) int {
	sort.Ints(nums)
	count := len(nums)

	for i := 0; i <= nums[len(nums)-1] && len(nums) >= i; i++ {
		for nums[0] < i {
			nums = nums[1:]
			count--
		}

		if i == count {
			return i
		}

	}

	return -1
}
