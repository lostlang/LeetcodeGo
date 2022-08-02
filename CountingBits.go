package leetcode

import "math/bits"

func countBits(n int) []int {
	var out = make([]int, n+1)

	for i := 0; i <= n; i++ {
		out[i] = bits.OnesCount(uint(i))
	}

	return out
}
