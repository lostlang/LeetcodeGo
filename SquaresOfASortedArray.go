package leetcode

import "sort"

func sortedSquares(nums []int) []int {
	var out = make([]int, len(nums))

	for i, val := range nums {
		out[i] = val * val
	}
	sort.Ints(out)
	return out
}
