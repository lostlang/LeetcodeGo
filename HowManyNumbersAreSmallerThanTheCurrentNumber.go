package leetcode

import "sort"

func smallerNumbersThanCurrent(nums []int) []int {

	var sortedCopy = make([]int, len(nums))
	copy(sortedCopy, nums)
	sort.Ints(sortedCopy)

	var hash = make(map[int]int)

	for i, val := range sortedCopy {
		if hash[val] == 0 {
			hash[val] = i + 1
		}
	}

	for i, val := range nums {
		sortedCopy[i] = hash[val] - 1
	}

	return sortedCopy
}
