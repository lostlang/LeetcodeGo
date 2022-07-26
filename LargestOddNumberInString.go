package leetcode

func largestOddNumber(num string) string {
	var numR = []rune(num)

	for len(numR) > 0 && numR[len(numR)-1] != '1' && numR[len(numR)-1] != '3' && numR[len(numR)-1] != '5' && numR[len(numR)-1] != '7' && numR[len(numR)-1] != '9' {
		numR = numR[:len(numR)-1]
	}

	return string(numR)
}
