package leetcode

func lengthOfLastWord(s string) int {
	var out int
	var tmp int

	for _, char := range s {
		if char == ' ' {
			if tmp > 0 {
				out = tmp
			}
			tmp = 0
		} else {
			tmp++
		}
	}

	if tmp > 0 {
		out = tmp
	}

	return out
}
