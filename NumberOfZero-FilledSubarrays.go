package leetcode

func zeroFilledSubarray(nums []int) int64 {
	var subZeros []int
	var maxLen int
	var out int64

	for _, val := range nums {
		if val == 0 {
			maxLen++
		} else {
			if maxLen != 0 {
				subZeros = append(subZeros, maxLen)
				maxLen = 0
			}
		}
	}

	if maxLen > 0 {
		subZeros = append(subZeros, maxLen)
	}

	for _, val := range subZeros {
		out += int64((val + 1) * val / 2)
	}

	return out
}
