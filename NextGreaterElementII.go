package leetcode

func nextGreaterElements(nums []int) []int {
	var double = make([]int, len(nums)*2)
	copy(double, nums)
	copy(double[len(nums):], nums)

	doubleNext := dailyTemperatures(double)

	for i := range nums {
		if doubleNext[i] == 0 {
			nums[i] = -1
		} else {
			nums[i] = double[i+doubleNext[i]]
		}
	}

	return nums
}
