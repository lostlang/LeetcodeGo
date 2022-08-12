package leetcode

import "unicode"

func letterCasePermutation(s string) []string {
	var out = []string{""}
	var cur []string

	for _, char := range s {
		cur = make([]string, 0)
		if unicode.IsLetter(char) {
			for _, str := range out {
				cur = append(cur, str+string(unicode.ToTitle(char)))
				cur = append(cur, str+string(unicode.ToLower(char)))
			}
		} else {
			for _, str := range out {
				cur = append(cur, str+string(char))
			}
		}

		out = cur
	}

	return out
}
