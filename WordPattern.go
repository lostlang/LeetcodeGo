package leetcode

import (
	"bytes"
)

func wordPattern(pattern string, s string) bool {
	var startChar = 'a'
	var hashPatter = make(map[rune]rune)
	var hashWords = make(map[string]rune)
	var newString, newWord bytes.Buffer

	for _, char := range s {
		if char == ' ' {
			if hashWords[newWord.String()] == 0 {
				hashWords[newWord.String()] = startChar
				startChar++
			}

			newString.WriteRune(hashWords[newWord.String()])
			newWord.Reset()
		} else {
			newWord.WriteRune(char)
		}
	}

	if hashWords[newWord.String()] == 0 {
		hashWords[newWord.String()] = startChar
		startChar++
	}

	newString.WriteRune(hashWords[newWord.String()])
	newWord.Reset()

	startChar = 'a'

	for _, char := range pattern {
		if hashPatter[char] == 0 {
			hashPatter[char] = startChar
			startChar++
		}

		newWord.WriteRune(hashPatter[char])
	}

	return newString.String() == newWord.String()
}
