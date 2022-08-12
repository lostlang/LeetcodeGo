package leetcode

import "math"

func judgeSquareSum(c int) bool {
	var i, j = 0, int(math.Sqrt(float64(c)))
	var sum int

	for i <= j {
		sum = i*i + j*j

		if sum == c {
			return true
		} else if sum < c {
			i++
		} else {
			j--
		}
	}

	return false
}
