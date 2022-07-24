package leetcode

import (
	"strings"
)

func firstUniqChar(s string) int {
	var hash = make(map[rune]int)

	for _, char := range s {
		hash[char]++
	}

	var minIndex int
	var out = -1

	for char, val := range hash {
		if val == 1 {
			minIndex = strings.IndexRune(s, char)
			if out == -1 || minIndex < out {
				out = minIndex
			}
		}
	}

	return out
}
