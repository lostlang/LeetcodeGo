package leetcode

func maxAreaOfIsland(grid [][]int) int {
	var out = 0

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				tmp := fillAdd(grid, i, j)

				if tmp > out {
					out = tmp
				}
			}
		}
	}

	return out
}

func fillAdd(grid [][]int, i int, j int) int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) || grid[i][j] != 1 {
		return 0
	}

	grid[i][j] = 0

	return 1 + fillAdd(grid, i-1, j) + fillAdd(grid, i+1, j) + fillAdd(grid, i, j-1) + fillAdd(grid, i, j+1)
}
