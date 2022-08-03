package leetcode

import "math/bits"

func countPrimeSetBits(left int, right int) int {
	var out int

	for i := left; i <= right; i++ {
		switch bits.OnesCount(uint(i)) {
		case 2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31:
			out++
		}
	}

	return out
}
