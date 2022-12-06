package leetcode

func spiralOrder(matrix [][]int) []int {
	var out []int

	for len(matrix) > 0 {
		out = append(out, matrix[0]...)
		matrix = matrix[1:]
		if len(matrix) == 0 {
			break
		}
		matrix = rotateMatrix(matrix)
	}

	return out
}

func rotateMatrix(matrix [][]int) [][]int {
	var newMatrix = make([][]int, len(matrix[0]))
	for i := range newMatrix {
		newMatrix[i] = make([]int, len(matrix))
	}

	for i := range matrix {
		for j := range matrix[i] {
			newMatrix[j][i] = matrix[i][len(matrix[i])-j-1]
		}
	}

	return newMatrix
}
