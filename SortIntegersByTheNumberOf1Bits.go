package leetcode

import (
	"math/bits"
)

func sortByBits(arr []int) []int {
	for i := range arr {
		for j := i + 1; j < len(arr); j++ {
			if !sortingBits(arr[i], arr[j]) {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

	return arr
}

func sortingBits(a, b int) bool {
	if bits.OnesCount(uint(a)) == bits.OnesCount(uint(b)) {
		return a < b
	}

	return bits.OnesCount(uint(a)) < bits.OnesCount(uint(b))
}
