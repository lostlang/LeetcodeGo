package leetcode

func missingNumber(nums []int) int {
	var out = len(nums) * (len(nums) + 1) / 2

	for _, num := range nums {
		out -= num
	}

	return out
}
