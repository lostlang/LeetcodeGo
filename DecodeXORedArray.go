package leetcode

func decode(encoded []int, first int) []int {
	var out = make([]int, len(encoded)+1)
	out[0] = first

	for i, val := range encoded {
		out[i+1] = out[i] ^ val
	}

	return out
}
