package leetcode

func isSubsequence(s string, t string) bool {
	var index int

	for index2 := range t {
		if index == len(s) {
			return true
		}
		if t[index2] == s[index] {
			index++
		}
	}

	return index == len(s)
}
