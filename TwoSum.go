package leetcode

import (
	"sort"
)

func twoSum(nums []int, target int) []int {
	var out = []int{-1, -1}
	var sortNums []int = make([]int, len(nums))
	copy(sortNums, nums)
	sort.Slice(sortNums, func(i, j int) bool {
		return sortNums[i] < sortNums[j]
	})

	var start = 0
	var end = len(sortNums) - 1

	var curSum = sortNums[start] + sortNums[end]

	for true {
		if curSum > target {
			curSum -= sortNums[end]
			end -= 1
			curSum += sortNums[end]
		}
		if curSum < target {
			curSum -= sortNums[start]
			start += 1
			curSum += sortNums[start]
		}
		if curSum == target {
			break
		}
	}
	start = sortNums[start]
	end = sortNums[end]
	for index, value := range nums {
		if start == value && out[0] == -1 {
			out[0] = index
		} else {
			if end == value && out[1] == -1 {
				out[1] = index
			}
		}
	}
	return out
}
