package leetcode

import "sort"

func removeDuplicates(nums []int) int {
	var dict = make(map[int]bool)
	for _, val := range nums {
		dict[val] = true
	}
	var array []int

	for val := range dict {
		array = append(array, val)
	}

	sort.Slice(array, func(i, j int) bool {
		return array[i] < array[j]
	})

	for index, val := range array {
		nums[index] = val
	}

	return len(dict)
}
