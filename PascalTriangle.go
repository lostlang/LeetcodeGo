package leetcode

func generate(numRows int) [][]int {
	var out = [][]int{[]int{1}}

	numRows--

	for numRows > 0 {
		var prev int
		var newLayer []int

		for _, val := range out[len(out)-1] {
			newLayer = append(newLayer, prev+val)
			prev = val
		}
		newLayer = append(newLayer, 1)

		out = append(out, newLayer)
		numRows--
	}

	return out
}
