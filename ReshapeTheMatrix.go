package leetcode

func matrixReshape(mat [][]int, r int, c int) [][]int {
	if len(mat)*len(mat[0]) != r*c {
		return mat
	}

	var out = make([][]int, r)

	for i := range out {
		out[i] = make([]int, c)
	}

	for i := 0; i < r*c; i++ {
		out[i/c][i%c] = mat[i/len(mat[0])][i%len(mat[0])]
	}

	return out
}
