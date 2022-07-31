package leetcode

import (
	"math"
)

func arrangeCoins(n int) int {
	var x = (math.Sqrt(float64(n*8+1)) - 1) / 2

	var out = int(x)
	out++

	return out - 1
}
