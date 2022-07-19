package leetcode

import (
	"bytes"
)

func isIsomorphic(s string, t string) bool {
	var sHash = make(map[rune]rune)
	var sNew bytes.Buffer
	var currenChar rune = 1

	for _, char := range s {
		if sHash[char] == 0 {
			sHash[char] = currenChar
			currenChar++
		}
		sNew.WriteRune(sHash[char])
	}

	sHash = make(map[rune]rune)
	var tNew bytes.Buffer
	currenChar = 1

	for _, char := range t {
		if sHash[char] == 0 {
			sHash[char] = currenChar
			currenChar++
		}
		tNew.WriteRune(sHash[char])
	}

	return sNew.String() == tNew.String()
}
