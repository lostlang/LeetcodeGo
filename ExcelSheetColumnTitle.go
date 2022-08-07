package leetcode

import (
	"bytes"
)

func convertToTitle(columnNumber int) string {
	var out bytes.Buffer

	for columnNumber > 0 {
		out.WriteRune(rune('A' + (columnNumber-1)%26))
		columnNumber = (columnNumber - 1) / 26
	}

	var tmp = []rune(out.String())
	out.Reset()

	for i := len(tmp) - 1; i >= 0; i-- {
		out.WriteRune(tmp[i])
	}

	return out.String()
}
