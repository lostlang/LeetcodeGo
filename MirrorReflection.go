package leetcode

func mirrorReflection(p int, q int) int {
	var newLen int

	for true {
		newLen += q

		if newLen%p == 0 {
			return newLen / p % 2
		}

		newLen += q

		if newLen%p == 0 && newLen/p%2 == 1 {
			return 2
		}
	}

	return 0
}
