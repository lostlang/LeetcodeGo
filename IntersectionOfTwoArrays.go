package leetcode

func intersection(nums1 []int, nums2 []int) []int {
	var hash1 = intToHash(nums1)
	var hash2 = intToHash(nums2)
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

func intToHash(arr []int) map[int]int {
	var out = make(map[int]int)

	for _, val := range arr {
		out[val]++
	}

	return out
}
