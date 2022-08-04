package leetcode

import "strings"

func findWords(words []string) []string {
	var out = []string{}
	var lover []rune
	var line int

Loop:
	for _, word := range words {
		lover = []rune(strings.ToLower(word))

		line = getLineChar(lover[0])

		for _, char := range lover {
			if line != getLineChar(char) {
				continue Loop
			}
		}

		out = append(out, word)
	}

	return out
}

func getLineChar(char rune) int {
	switch char {
	case 'q', 'w', 'e', 'r', 't', 'y', 'u', 'i', 'o', 'p':
		return 1
	case 'a', 's', 'd', 'f', 'g', 'h', 'j', 'k', 'l':
		return 2
	case 'z', 'x', 'c', 'v', 'b', 'n', 'm':
		return 3
	}

	return 0
}
