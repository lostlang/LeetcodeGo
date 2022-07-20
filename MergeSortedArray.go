package leetcode

import "sort"

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if nums1[i] > nums2[j] {
				nums1[i], nums2[j] = nums2[j], nums1[i]
			}
		}
	}

	sort.Ints(nums2)

	for i := 0; i < n; i++ {
		nums1[i+m] = nums2[i]
	}
}
