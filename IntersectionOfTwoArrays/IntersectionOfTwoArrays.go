package IntersectionOfTwoArrays

import "leetcode/utils"

var IntToHash = utils.IntToHash

func intersection(nums1 []int, nums2 []int) []int {
	hash1 := IntToHash(nums1)
	hash2 := IntToHash(nums2)
	var out []int
	var min int

	for val := range hash1 {
		min = hash1[val]
		if min > hash2[val] {
			min = hash2[val]
		}

		if min > 0 {
			out = append(out, val)
		}
	}

	return out
}
