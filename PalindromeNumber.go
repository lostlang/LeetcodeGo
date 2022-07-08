package leetcode

func isPalindromeInt(x int) bool {
	if x < 0 {
		return false
	}

	var maxX = 1
	for x/maxX >= 10 {
		maxX *= 10
	}

	var first, last int
	for x > 0 {
		first = x / maxX
		last = x % 10
		if first != last {
			return false
		}
		x -= first * maxX
		x /= 10
		maxX /= 100
	}

	return true
}
