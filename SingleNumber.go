package leetcode

func singleNumber(nums []int) int {
	var hash = intToHash(nums)

	for i, val := range hash {
		if val == 1 {
			return i
		}
	}

	return 0
}
