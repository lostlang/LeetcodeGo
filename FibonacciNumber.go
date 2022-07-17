package leetcode

func fib(n int) int {
	var num1 int
	var num2 = 1

	if n <= 1 {
		return n
	}

	for i := 0; i < n; i++ {
		num1, num2 = num2, num1+num2
	}

	return num1
}
