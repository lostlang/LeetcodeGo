package leetcode

func isPowerOfTwo(n int) bool {
	var out = 1

	for out < n {
		out *= 2
	}

	if out == n {
		return true
	}

	return false
}
