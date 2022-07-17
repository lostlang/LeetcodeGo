package leetcode

import (
	"sort"
)

func minimumSum(num int) int {
	var numAr = make([]int, 4)

	for i := 0; i < 4; i++ {
		numAr[i] = num % 10
		num /= 10
	}

	sort.Ints(numAr)

	var sum = (numAr[0]+numAr[1])*10 + numAr[2] + numAr[3]

	return sum
}
