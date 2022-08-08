package leetcode

func findComplement(num int) int {
	var out int
	var bit int

	for num > 0 {
		out |= (num&1 ^ 1) << bit
		num >>= 1
		bit += 1
	}

	return out
}
