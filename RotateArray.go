package leetcode

func rotate(nums []int, k int) {
	k = k % len(nums)

	var tmp = append(nums[len(nums)-k:], nums[:len(nums)-k]...)

	for i := range nums {
		nums[i] = tmp[i]
	}

	return
}
