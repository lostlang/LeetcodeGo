package leetcode

func countDigits(num int) int {
	output := 0
	digits := numberToDigits(num)

	for _, digit := range digits {
		if digit != 0 && num%digit == 0 {
			output++
		}
	}

	return output
}

func numberToDigits(num int) []int {
	output := []int{}

	for num > 0 {
		output = append(output, num%10)
		num /= 10
	}

	return output
}
