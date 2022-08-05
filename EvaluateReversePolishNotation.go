package leetcode

import "strconv"

func evalRPN(tokens []string) int {
	var intTokens []int
	var tmp int

	for _, val := range tokens {
		switch val {
		case "+":
			intTokens[len(intTokens)-2] = intTokens[len(intTokens)-2] + intTokens[len(intTokens)-1]
			intTokens = intTokens[:len(intTokens)-1]
		case "-":
			intTokens[len(intTokens)-2] = intTokens[len(intTokens)-2] - intTokens[len(intTokens)-1]
			intTokens = intTokens[:len(intTokens)-1]
		case "*":
			intTokens[len(intTokens)-2] = intTokens[len(intTokens)-2] * intTokens[len(intTokens)-1]
			intTokens = intTokens[:len(intTokens)-1]
		case "/":
			intTokens[len(intTokens)-2] = intTokens[len(intTokens)-2] / intTokens[len(intTokens)-1]
			intTokens = intTokens[:len(intTokens)-1]
		default:
			tmp, _ = strconv.Atoi(val)
			intTokens = append(intTokens, tmp)
		}
	}

	return intTokens[0]
}
