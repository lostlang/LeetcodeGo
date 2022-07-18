package leetcode

import "strings"

func mostWordsFound(sentences []string) int {
	var out int

	for _, str := range sentences {
		var cur = strings.Count(str, " ")
		if cur > out {
			out = cur
		}
	}

	out++

	return out
}
