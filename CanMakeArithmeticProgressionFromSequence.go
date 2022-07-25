package leetcode

import "sort"

func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)

	var dif = arr[1] - arr[0]

	for i := 1; i < len(arr); i++ {
		if dif != arr[i]-arr[i-1] {
			return false
		}
	}

	return true
}
