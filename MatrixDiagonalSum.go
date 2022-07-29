package leetcode

func diagonalSum(mat [][]int) int {
	var sum = 0

	if len(mat) == 0 {
		return sum
	}

	for i := range mat {
		sum += mat[i][i]
		if len(mat)-i-1 != i {
			sum += mat[i][len(mat)-i-1]
		}
	}

	return sum
}
