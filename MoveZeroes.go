package leetcode

func moveZeroes(nums []int) {
	var index int

	for _, val2 := range nums {
		nums[index] = val2
		if val2 != 0 {
			index++
		}
	}

	for ; index < len(nums); index++ {
		nums[index] = 0
	}
}
