package leetcode

func getConcatenation(nums []int) []int {
	var out = make([]int, len(nums)*2)
	copy(out, nums)
	for index, value := range nums {
		out[index+len(nums)] = value
	}
	return out
}
