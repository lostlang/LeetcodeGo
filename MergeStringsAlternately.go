package leetcode

import (
	"bytes"
)

func mergeAlternately(word1 string, word2 string) string {
	var out bytes.Buffer
	var minMaxLen int
	var rW1 = []rune(word1)
	var rW2 = []rune(word2)

	if len(rW1) > len(rW2) {
		minMaxLen = len(rW2)
	} else {
		minMaxLen = len(rW1)
	}

	for i := 0; i < minMaxLen; i++ {
		out.WriteRune(rW1[i])
		out.WriteRune(rW2[i])
	}

	if len(rW1) > minMaxLen {
		for i := minMaxLen; i < len(rW1); i++ {
			out.WriteRune(rW1[i])
		}

	}

	if len(rW2) > minMaxLen {
		for i := minMaxLen; i < len(rW2); i++ {
			out.WriteRune(rW2[i])
		}
	}

	return out.String()
}
