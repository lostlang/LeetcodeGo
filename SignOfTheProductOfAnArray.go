package leetcode

func arraySign(nums []int) int {
	var out = 1
	for _, val := range nums {
		if val == 0 {
			return 0
		}

		if val < 0 {
			out *= -1
		}
	}

	return out
}
