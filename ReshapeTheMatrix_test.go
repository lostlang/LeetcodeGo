package leetcode

import (
	"reflect"
	"testing"
)

type matrixReshapeTestPair struct {
	inputArr    [][]int
	inputRow    int
	inputColumn int
	out         [][]int
}

var matrixReshapeTestCases = []matrixReshapeTestPair{
	{[][]int{[]int{1, 2}, []int{3, 4}}, 1, 4, [][]int{[]int{1, 2, 3, 4}}},
	{[][]int{[]int{1, 2}, []int{3, 4}}, 2, 4, [][]int{[]int{1, 2}, []int{3, 4}}},
}

func TestEvalMatrixReshape(t *testing.T) {
	for _, pair := range matrixReshapeTestCases {
		var newOut = matrixReshape(pair.inputArr, pair.inputRow, pair.inputColumn)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputArr,
				",", pair.inputRow,
				"and", pair.inputColumn,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
