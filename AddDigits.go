package leetcode

func addDigits(num int) int {
	if num < 10 {
		return num
	}

	var sum = 0

	for num > 0 {
		sum += num % 10
		num /= 10
	}

	return addDigits(sum)
}
