package leetcode

func hasAlternatingBits(n int) bool {
	var i = 1
	var swap bool

	for i < n {
		if swap {
			i = 2*i + 1
		} else {
			i *= 2
		}

		swap = !swap
	}

	return i == n
}
