package leetcode

import (
	"fmt"
)

func minSumSquareDiff(nums1 []int, nums2 []int, k1 int, k2 int) int64 {
	var out = make([]int, len(nums1))
	var k = k1 + k2
	var minOut int

	if nums1[0] > nums2[0] {
		minOut = nums1[0] - nums2[0]
	} else {
		minOut = nums2[0] - nums1[0]
	}

	for index := range nums1 {
		out[index] = nums1[index] - nums2[index]
		if nums1[index] < nums2[index] {
			out[index] = -out[index]
		}

	}

	var sum int64

	for _, val := range out {
		sum += int64(val) * int64(val)
	}

	fmt.Println(k, minOut, out)

	return sum
}
