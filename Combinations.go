package leetcode

func combine(n int, k int) [][]int {
	return combineRecursive(1, n, k)
}

func combineRecursive(min, max, k int) [][]int {
	var out [][]int

	if k == 1 {
		for i := min; i <= max; i++ {
			out = append(out, []int{i})
		}

		return out
	}

	for i := min; i < max; i++ {
		for _, line := range combineRecursive(i+1, max, k-1) {
			out = append(out, append([]int{i}, line...))
		}
	}

	return out
}
