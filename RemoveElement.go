package leetcode

func removeElement(nums []int, val int) int {
	var index int

	for _, val2 := range nums {
		nums[index] = val2
		if val2 != val {
			index++
		}
	}

	return index
}
