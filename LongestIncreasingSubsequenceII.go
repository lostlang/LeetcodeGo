package leetcode

func lengthOfLIS2(nums []int, k int) int {
	var out int
	var dp = make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		for j := i; j >= 0; j-- {
			if nums[i] > nums[j] && nums[i]-nums[j] <= k && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}

			if j <= dp[i]-1 {
				break
			}
		}

		if dp[i] == 0 {
			dp[i] = 1
		}
	}

	for i := 0; i < len(nums); i++ {
		if dp[i] > out {
			out = dp[i]
		}
	}

	return out
}
