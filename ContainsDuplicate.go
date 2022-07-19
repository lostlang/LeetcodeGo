package leetcode

func containsDuplicate(nums []int) bool {
	var hashTable = make(map[int]bool)

	for _, val := range nums {
		if hashTable[val] {
			return true
		} else {
			hashTable[val] = true
		}
	}

	return false
}
