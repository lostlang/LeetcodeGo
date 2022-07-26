package leetcode

func isSameAfterReversals(num int) bool {
	if num == 0 {
		return true
	}

	if num%10 == 0 {
		return false
	}

	return true
}
