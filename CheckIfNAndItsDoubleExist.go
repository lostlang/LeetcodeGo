package leetcode

func checkIfExist(arr []int) bool {
	var hash = make(map[int]bool)
	var count int

	for _, val := range arr {
		hash[val] = true
		if val == 0 {
			count++
		}
	}

	for val := range hash {
		if hash[2*val] && val != 0 {
			return true
		}

		if val == 0 && count > 1 {
			return true
		}
	}

	return false
}
