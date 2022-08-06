package leetcode

import (
	"bytes"
	"unicode"
)

func isPalindromeStr(s string) bool {
	var palindrome bytes.Buffer

	for _, char := range s {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			palindrome.WriteRune(unicode.ToLower(char))
		}
	}

	var st = palindrome.String()

	for i := 0; i < len(st)/2; i++ {
		if st[i] != st[len(st)-1-i] {
			return false
		}
	}

	return true
}
