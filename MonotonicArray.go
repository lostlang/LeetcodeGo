package leetcode

func isMonotonic(nums []int) bool {
	var sumDif, sumSame int

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			sumSame++
		} else if nums[i] > nums[i-1] {
			sumDif++
		} else {
			sumDif--
		}
	}

	return sumDif+sumSame == len(nums)-1 || -sumDif+sumSame == len(nums)-1
}
