package leetcode

import "bytes"

func interpret(command string) string {
	var flag bool
	var out bytes.Buffer

	for _, char := range command {
		switch char {
		case '(':
			flag = true
		case ')':
			if flag {
				out.WriteRune('o')
				flag = false
			}
		case 'a':
			if flag {
				out.WriteString("al")
				flag = false
			}
		case 'G':
			out.WriteRune('G')
		}
	}

	return out.String()
}
