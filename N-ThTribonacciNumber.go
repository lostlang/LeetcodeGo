package leetcode

func tribonacci(n int) int {
	var num1 int
	var num2, num3 = 1, 1

	if n <= 1 {
		return n
	}
	if n == 2 {
		return 1
	}

	for i := 0; i < n; i++ {
		num1, num2, num3 = num2, num3, num1+num2+num3
	}

	return num1
}
