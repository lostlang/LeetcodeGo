package leetcode

func isPerfectSquare(num int) bool {
	var out = 0

	for out*out < num {
		out++
	}

	if out*out == num {
		return true
	}

	return false
}
