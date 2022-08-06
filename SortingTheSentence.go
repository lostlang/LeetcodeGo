package leetcode

import (
	"bytes"
)

func sortSentence(s string) string {
	var out = ""
	var words = make(map[rune]string)
	var tmp bytes.Buffer
	var wordCount int

	for _, c := range s {
		switch {
		case c >= '1' && c <= '9':
			words[c] = tmp.String()
			tmp.Reset()
			wordCount++
		case c != ' ':
			tmp.WriteRune(c)
		}
	}

	for i := 0; i < wordCount; i++ {
		out = out + words[rune(i+'1')]
		if i < wordCount-1 {
			out = out + " "
		}
	}

	return out
}
