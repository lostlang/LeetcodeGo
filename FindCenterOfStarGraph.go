package leetcode

func findCenter(edges [][]int) int {
	output := 0
	node := make([]int, 4)
	node[0] = edges[0][0]
	node[1] = edges[0][1]
	node[2] = edges[1][0]
	node[3] = edges[1][1]

	for i := 0; i < 4; i++ {
		for j := i + 1; j < 4; j++ {
			if node[i] == node[j] {
				output = node[i]
			}
		}
	}

	return output
}
