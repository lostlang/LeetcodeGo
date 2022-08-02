package leetcode

import "sort"

func kthSmallest(matrix [][]int, k int) int {
	var toLine = make([]int, len(matrix)*len(matrix[0]))

	for i, line := range matrix {
		copy(toLine[i*len(matrix[0]):(i+1)*len(matrix[0])], line)
	}

	sort.Ints(toLine)

	return toLine[k-1]
}
