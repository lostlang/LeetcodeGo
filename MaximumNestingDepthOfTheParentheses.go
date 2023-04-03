package leetcode

func maxDepthPairs(s string) int {
	output := 0
	depth := 0
	for _, c := range s {
		if c == '(' {
			depth++
		}
		if c == ')' {
			depth--
		}

		if depth > output {
			output = depth
		}
	}

	return output
}
