package leetcode

import (
	"sort"
)

func nextGreaterElementIII(n int) int {
	var digits = make([]int, 0)
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}

	var flag bool

	for i := 1; i < len(digits); i++ {
		if digits[i] < digits[i-1] {
			flag = true

			sort.Ints(digits[:i])
			index := searchInsert(digits[:i], digits[i])

			for digits[index] == digits[i] {
				index++
			}

			digits[i], digits[index] = digits[index], digits[i]

			sort.Ints(digits[:i])

			for j := 0; j < i/2; j++ {
				digits[j], digits[i-j-1] = digits[i-j-1], digits[j]
			}

			break
		}
	}

	if !flag {
		return -1
	}

	for i := len(digits) - 1; i >= 0; i-- {
		n *= 10
		n += digits[i]
	}

	if n <= 2_147_483_647 {
		return n
	}

	return -1
}
