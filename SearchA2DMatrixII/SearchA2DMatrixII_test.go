package SearchA2DMatrixII

import (
	"reflect"
	"testing"
)

type searchMatrixTestPair struct {
	inputMatrix [][]int
	inputTarget int
	output      bool
}

var searchMatrixTestCases = []searchMatrixTestPair{
	{[][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 5, true},
	{[][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 20, false},
}

func TestEvalSearchMatrix(t *testing.T) {
	for _, pair := range searchMatrixTestCases {
		newOutput := searchMatrix(pair.inputMatrix, pair.inputTarget)
		if !reflect.DeepEqual(newOutput, pair.output) {
			t.Error(
				"For", pair.inputMatrix, pair.inputTarget,
				"expected", pair.output,
				"got", newOutput,
			)
		}
	}
}
