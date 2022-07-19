package leetcode

func getRow(rowIndex int) []int {
	var out = [][]int{[]int{1}}

	for rowIndex > 0 {
		var prev int
		var newLayer []int

		for _, val := range out[len(out)-1] {
			newLayer = append(newLayer, prev+val)
			prev = val
		}
		newLayer = append(newLayer, 1)

		out = append(out, newLayer)
		rowIndex--
	}

	return out[len(out)-1]
}
