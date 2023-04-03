package leetcode

import (
	"strings"
)

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	string1 := strings.Join(word1, "")
	string2 := strings.Join(word2, "")

	if string1 == string2 {
		return true
	}

	return false
}
