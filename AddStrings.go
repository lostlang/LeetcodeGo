package leetcode

import (
	"bytes"
)

func addStrings(num1 string, num2 string) string {
	var array1 = stringToArrayInts(num1)
	var array2 = stringToArrayInts(num2)

	var sum = sumArrays(array1, array2)
	var out bytes.Buffer

	for _, digit := range sum {
		out.WriteRune(rune(digit + '0'))
	}

	return out.String()
}

func stringToArrayInts(s string) []int {
	var out []int

	for _, char := range s {
		out = append(out, int(char-'0'))
	}

	return out
}

func sumArrays(a []int, b []int) []int {
	var out []int
	var carry int
	var i = len(a) - 1
	var j = len(b) - 1

	for i >= 0 || j >= 0 {
		var sum int

		if i >= 0 {
			sum += a[i]
		}

		if j >= 0 {
			sum += b[j]
		}

		sum += carry

		if sum >= 10 {
			carry = 1
			sum -= 10
		} else {
			carry = 0
		}
		out = append(out, sum)

		i--
		j--
	}

	if carry > 0 {
		out = append(out, carry)
	}

	return reverseInts(out)
}

func reverseInts(a []int) []int {
	var out []int

	for i := len(a) - 1; i >= 0; i-- {
		out = append(out, a[i])
	}

	return out
}
