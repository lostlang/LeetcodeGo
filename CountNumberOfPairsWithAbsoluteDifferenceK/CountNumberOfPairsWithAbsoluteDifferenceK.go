package CountNumberOfPairsWithAbsoluteDifferenceK

import "leetcode/utils"

var IntToHash = utils.IntToHash

func countKDifference(nums []int, k int) int {
	var out int
	hash := IntToHash(nums)

	for val := range hash {
		out += hash[val] * (hash[val+k])
	}

	return out
}
