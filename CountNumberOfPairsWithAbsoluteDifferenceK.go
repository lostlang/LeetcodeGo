package leetcode

func countKDifference(nums []int, k int) int {
	var out int
	var hash = intToHash(nums)

	for val := range hash {
		out += hash[val] * (hash[val+k])
	}

	return out
}
