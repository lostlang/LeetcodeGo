package leetcode

func leftRigthDifference(nums []int) []int {
	out := make([]int, len(nums))
	leftSum := 0
	rightSum := 0
	for _, v := range nums {
		rightSum += v
	}

	for i := range nums {
		rightSum -= nums[i]
		if leftSum > rightSum {
			out[i] = leftSum - rightSum
		} else {
			out[i] = rightSum - leftSum
		}
		leftSum += nums[i]
	}

	return out
}
