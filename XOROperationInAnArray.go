package leetcode

func xorOperation(n int, start int) int {
	var out int

	for i := 0; i < n; i++ {
		out ^= start + 2*i
	}

	return out
}
