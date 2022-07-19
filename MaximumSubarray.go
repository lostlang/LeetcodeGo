package leetcode

func maxSubArray(nums []int) int {
	var maxSum = nums[0]
	var curSum int

	for _, val := range nums {
		curSum += val

		if curSum > maxSum {
			maxSum = curSum
		}

		if curSum < 0 {
			curSum = 0
		}
	}

	return maxSum
}
