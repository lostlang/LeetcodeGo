package leetcode

func islandPerimeter(grid [][]int) int {
	var out int

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				out += 4
				if i > 0 && grid[i-1][j] == 1 {
					out -= 2
				}
				if j > 0 && grid[i][j-1] == 1 {
					out -= 2
				}
			}
		}
	}

	return out
}
