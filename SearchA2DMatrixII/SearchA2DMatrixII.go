package SearchA2DMatrixII

import (
	"sort"

	"leetcode/BinarySearch"
)

var search = BinarySearch.Search

func searchMatrix(matrix [][]int, target int) bool {
	if search(toFlat(matrix), target) != -1 {
		return true
	}
	return false
}

func toFlat(matrix [][]int) []int {
	h, w := len(matrix), len(matrix[0])
	out := make([]int, h*w)

	for i := 0; i < h*w; i++ {
		out[i] = matrix[i/w][i%w]
	}

	sort.Ints(out)

	return out
}
