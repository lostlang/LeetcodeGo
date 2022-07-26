package leetcode

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var out = make([]int, len(nums1))
	var hash = make(map[int]int)

	for i, val := range nums2 {
		hash[val] = nextMax(val, nums2[i:])
	}

	for i, val := range nums1 {
		out[i] = hash[val]
	}

	return out
}

func nextMax(val int, arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	for _, val2 := range arr {
		if val2 > val {
			return val2
		}
	}

	return -1
}
