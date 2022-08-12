package leetcode

func numIslands(grid [][]byte) int {
	var out int

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				out++
				fill(grid, i, j, 1)
			}
		}
	}

	return out
}

func fill(grid [][]byte, i int, j int, c byte) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) || grid[i][j] != c {
		return
	}

	grid[i][j] = '0'
	fill(grid, i-1, j, c)
	fill(grid, i+1, j, c)
	fill(grid, i, j-1, c)
	fill(grid, i, j+1, c)
}
