package leetcode

func buildArray(nums []int) []int {
	var out = make([]int, len(nums))
	for index, value := range nums {
		out[index] = nums[value]
	}
	return out
}
