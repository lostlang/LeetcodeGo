package leetcode

import "strconv"

func fizzBuzz(n int) []string {
	var out = make([]string, n)
	for index := range out {
		switch (index + 1) % 15 {
		case 0:
			out[index] = "FizzBuzz"
		case 3, 6, 9, 12:
			out[index] = "Fizz"
		case 5, 10:
			out[index] = "Buzz"
		default:
			out[index] = strconv.Itoa(index + 1)
		}
	}
	return out
}
