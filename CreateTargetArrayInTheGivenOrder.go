package leetcode

func createTargetArray(nums []int, index []int) []int {
	var out = make([]int, 0)

	for i, v := range index {
		out = append(out[:v], append([]int{nums[i]}, out[v:]...)...)
	}

	return out
}
