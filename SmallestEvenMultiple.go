package leetcode

func smallestEvenMultiple(n int) int {
	var out int

	if n%2 == 0 {
		out = n
	} else {
		out = n * 2
	}

	return out
}
