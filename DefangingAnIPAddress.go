package leetcode

import "bytes"

func defangIPaddr(address string) string {
	var out bytes.Buffer

	for _, char := range address {

		if char != '.' {
			out.WriteRune(char)
		} else {
			out.WriteString("[.]")
		}

	}

	return out.String()
}
