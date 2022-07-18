package leetcode

func numberOfSteps(num int) int {
	var out int
	for num > 0 {
		if num%2 == 0 {
			num /= 2
		} else {
			num--
		}
		out++
	}

	return out
}
