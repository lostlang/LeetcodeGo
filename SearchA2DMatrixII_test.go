package leetcode

import (
	"reflect"
	"testing"
)

type searchMatrixTestPair struct {
	inputMatrix [][]int
	inputTarget int
	out         bool
}

var searchMatrixTestCases = []searchMatrixTestPair{
	{[][]int{[]int{1, 4, 7, 11, 15}, []int{2, 5, 8, 12, 19}, []int{3, 6, 9, 16, 22}, []int{10, 13, 14, 17, 24}, []int{18, 21, 23, 26, 30}}, 5, true},
	{[][]int{[]int{1, 4, 7, 11, 15}, []int{2, 5, 8, 12, 19}, []int{3, 6, 9, 16, 22}, []int{10, 13, 14, 17, 24}, []int{18, 21, 23, 26, 30}}, 20, false},
}

func TestEvalSearchMatrix(t *testing.T) {
	for _, pair := range searchMatrixTestCases {
		var newOut = searchMatrix(pair.inputMatrix, pair.inputTarget)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputMatrix,
				"and", pair.inputTarget,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
