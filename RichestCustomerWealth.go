package leetcode

func maximumWealth(accounts [][]int) int {
	var sum int

	var curSum int

	for _, arr := range accounts {
		for _, val := range arr {
			curSum += val
		}
		if curSum > sum {
			sum = curSum
		}

		curSum = 0
	}

	return sum
}
