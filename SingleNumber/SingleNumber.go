package SingleNumber

import "leetcode/utils"

var IntToHash = utils.IntToHash

func singleNumber(nums []int) int {
	hash := IntToHash(nums)

	for i, val := range hash {
		if val == 1 {
			return i
		}
	}

	return 0
}
