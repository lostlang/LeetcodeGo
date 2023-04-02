package leetcode

func differenceOfSum(nums []int) int {
	output := 0
	sum := 0
	sumOfDigits := 0
	for _, num := range nums {
		sum += num
		sumOfDigits += digitSum(num)
	}

	if sum > sumOfDigits {
		output = sum - sumOfDigits
	} else {
		output = sumOfDigits - sum
	}

	return output
}

func digitSum(n int) int {
	output := 0

	for n > 0 {
		output += n % 10
		n /= 10
	}

	return output
}
