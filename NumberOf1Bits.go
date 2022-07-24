package leetcode

func hammingWeight(num uint32) int {
	var out int

	for num > 0 {
		if num%2 == 1 {
			out++
		}

		num /= 2
	}

	return out
}
