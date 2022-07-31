package leetcode

import (
	"bytes"
	"strconv"
)

func freqAlphabets(s string) string {
	var out bytes.Buffer
	var tmp bytes.Buffer

	for _, char := range s {
		if char == '#' {
			out.WriteString(convertNumToAlphabet(tmp, true))
			tmp.Reset()
		} else {
			tmp.WriteRune(char)
		}
	}

	if tmp.Len() > 0 {
		out.WriteString(convertNumToAlphabet(tmp, false))
	}

	return out.String()
}

func convertNumToAlphabet(buf bytes.Buffer, split bool) string {
	var out bytes.Buffer
	var val int

	for buf.Len() > 2 || !split && buf.Len() > 0 {
		val, _ = strconv.Atoi(string(buf.Next(1)))
		out.WriteByte(byte(val + 96))
	}

	if buf.Len() > 0 {
		val, _ = strconv.Atoi(buf.String())
		out.WriteByte(byte(val + 96))
	}

	return out.String()
}
