package leetcode

import "strings"

func truncateSentence(s string, k int) string {
	output := ""
	words := strings.Split(s, " ")
	output = strings.Join(words[:k], " ")

	return output
}
