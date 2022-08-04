package leetcode

import (
	"reflect"
	"testing"
)

type isToeplitzMatrixTestPair struct {
	input [][]int
	out   bool
}

var isToeplitzMatrixTestCases = []isToeplitzMatrixTestPair{
	{[][]int{[]int{1, 2, 3, 4}, []int{5, 1, 2, 3}, []int{9, 5, 1, 2}}, true},
	{[][]int{[]int{1, 2}, []int{2, 2}}, false},
}

func TestEvalIsToeplitzMatrix(t *testing.T) {
	for _, pair := range isToeplitzMatrixTestCases {
		var newOut = isToeplitzMatrix(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
