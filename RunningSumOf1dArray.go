package leetcode

func runningSum(nums []int) []int {
	var out = make([]int, len(nums))
	var sum int
	for index, value := range nums {
		sum += value
		out[index] = sum
	}
	return out
}
