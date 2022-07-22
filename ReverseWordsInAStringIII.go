package leetcode

import (
	"bytes"
)

func reverseWords(s string) string {
	var newS bytes.Buffer
	var word bytes.Buffer

	for _, char := range s {
		if char != ' ' {
			word.WriteRune(char)
		} else {
			reverseAdd(&word, &newS)
			newS.WriteRune(' ')
		}
	}

	reverseAdd(&word, &newS)

	return newS.String()
}

func reverseAdd(word, newString *bytes.Buffer) {
	var wordString []rune

	wordString = []rune(word.String())

	for i := len(wordString) - 1; i >= 0; i-- {
		newString.WriteRune(wordString[i])
	}

	word.Reset()
}
