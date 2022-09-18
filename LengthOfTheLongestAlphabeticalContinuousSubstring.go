package leetcode

func longestContinuousSubstring(s string) int {
	var out int
	var current = 1

	for i, char := range s {
		if i > 0 && char == rune(s[i-1]+1) {
			current++
		} else {
			current = 1
		}

		if current > out {
			out = current
		}
	}

	return out
}
