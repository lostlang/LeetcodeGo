package leetcode

func decompressRLElist(nums []int) []int {
	var out, tmp []int

	for i := 0; i < len(nums); i += 2 {
		tmp = make([]int, nums[i])

		for j := range tmp {
			tmp[j] = nums[i+1]
		}

		out = append(out, tmp...)
	}

	return out
}
