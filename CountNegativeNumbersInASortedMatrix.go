package leetcode

func countNegatives(grid [][]int) int {
	var out int
	var start = len(grid[0]) - 1

	for _, line := range grid {
		for start >= 0 && line[start] < 0 {
			start--
		}

		out += len(line) - start - 1
	}

	return out
}
