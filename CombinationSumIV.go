package leetcode

import (
	"sort"
)

func combinationSum4(nums []int, target int) int {
	sort.Ints(nums)
	var hash = make(map[int]int)
	hash[0] = 1

	return hashedSum(nums, target, hash)
}

func hashedSum(nums []int, target int, hash map[int]int) int {
	if target < 0 {
		return 0
	}

	if hash[target] == -1 {
		return 0
	}

	if hash[target] != 0 {
		return hash[target]
	}

	var out int
	var maxIndex = searchInsert(nums, target) + 1
	if maxIndex >= len(nums) {
		maxIndex = len(nums) - 1
	}

	for i := 0; i <= maxIndex; i++ {
		out += hashedSum(nums, target-nums[i], hash)
	}

	if out == 0 {
		hash[target] = -1
	} else {
		hash[target] = out
	}

	return out
}
