package BuildArrayFromPermutation

func buildArray(nums []int) []int {
	out := make([]int, len(nums))
	for index, value := range nums {
		out[index] = nums[value]
	}
	return out
}
