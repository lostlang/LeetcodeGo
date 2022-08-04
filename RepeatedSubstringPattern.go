package leetcode

import (
	"bytes"
	"strings"
)

func repeatedSubstringPattern(s string) bool {
	var buf bytes.Buffer
	var rS = []rune(s)
	var tmp string

	for i, char := range rS {
		buf.WriteRune(char)

		if len(rS)%(i+1) == 0 {
			tmp = strings.Repeat(buf.String(), len(rS)/(i+1))

			if tmp == s && i != len(rS)-1 {
				return true
			}
		}
	}

	return false
}
