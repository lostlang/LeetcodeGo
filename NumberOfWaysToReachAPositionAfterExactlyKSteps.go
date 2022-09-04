package leetcode

func numberOfWays(startPos int, endPos int, k int) int {
	if (startPos+k)%2 != endPos%2 {
		return 0
	}

	var dp = make(map[int]int)
	dp[startPos] = 0
	var minPos = startPos
	var maxPos = startPos

	for ; k > 0; k-- {
		minPos--
		maxPos++

		dp2 := make(map[int]int)

		for pos := range dp {
			dp2[pos] = (dp[pos-1] + dp[pos+1]) % 1000000007
		}

		dp = dp2

		dp[minPos] = 1
		dp[maxPos] = 1

	}

	return dp[endPos]
}
