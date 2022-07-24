package leetcode

import (
	"sort"
)

func largestPerimeter(nums []int) int {
	var out int
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if nums[i+2] < nums[i+1]+nums[i] {
			out = nums[i] + nums[i+1] + nums[i+2]
		}
	}

	return out
}
