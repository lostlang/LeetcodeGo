package leetcode

import (
	"sort"
)

func makesquare(matchsticks []int) bool {

	var sum = 0
	for _, val := range matchsticks {
		sum += val
	}

	if sum%4 != 0 {
		return false
	}

	var avg = sum / 4
	sort.Ints(matchsticks)
	var sides = make([]int, 4)

	return backtrack(0, avg, matchsticks, &sides)

}

func backtrack(i, sideLen int, matchsticks []int, sides *[]int) bool {
	if i == len(matchsticks) {
		return true
	}

	for j := range *sides {
		if (*sides)[j]+matchsticks[i] <= sideLen {
			(*sides)[j] += matchsticks[i]
			if backtrack(i+1, sideLen, matchsticks, sides) {
				return true
			}
			(*sides)[j] -= matchsticks[i]
		}
	}

	return false
}
