package leetcode

func lengthOfLongestSubstring(s string) int {

	if len(s) < 2 {
		return len(s)
	}

	var out = 0
	var start = 0
	var end = 0

	for end < len(s) {
		end++

		start += shift(s[start:end])

		if end-start > out {
			out = end - start
		}
	}

	return out
}

func shift(s string) int {
	if len(s) < 2 {
		return 0
	}

	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[len(s)-1] {
			return i + 1
		}
	}

	return 0
}
