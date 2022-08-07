package leetcode

import (
	"bytes"
)

func decodeMessage(key string, message string) string {
	var out bytes.Buffer
	var hash = make(map[rune]rune)
	var i int

	for _, char := range key {
		if char != ' ' && hash[char] == 0 {
			hash[char] = 'a' + rune(i)
			i++
		}
	}

	for _, char := range message {
		if char == ' ' {
			out.WriteRune(' ')
		} else {
			out.WriteRune(hash[char])
		}
	}

	return out.String()
}
