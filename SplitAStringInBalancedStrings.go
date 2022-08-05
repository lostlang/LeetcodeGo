package leetcode

func balancedStringSplit(s string) int {
	var out, sum int

	for _, char := range s {
		switch char {
		case 'R':
			sum++
		case 'L':
			sum--
		}

		if sum == 0 {
			out++
		}
	}

	return out
}
