package leetcode

func subtractProductAndSum(n int) int {
	var sum int
	var prod = 1

	for n > 0 {
		sum += n % 10
		prod *= n % 10
		n /= 10
	}

	return prod - sum
}
