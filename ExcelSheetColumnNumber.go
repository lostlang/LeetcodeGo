package leetcode

func titleToNumber(columnTitle string) int {
	var out int

	for _, char := range columnTitle {
		out = out*26 + int(char-'A') + 1
	}

	return out
}
