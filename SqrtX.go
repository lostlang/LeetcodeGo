package leetcode

func mySqrt(x int) int {
	var out = 0

	for out*out < x {
		out++
	}

	if out*out > x {
		out--
	}

	return out
}
