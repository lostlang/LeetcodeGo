package leetcode

import (
	"reflect"
	"testing"
)

type diagonalSumTestPair struct {
	input [][]int
	out   int
}

var diagonalSumTestCases = []diagonalSumTestPair{
	{[][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}, 25},
	{[][]int{[]int{1, 1, 1, 1}, []int{1, 1, 1, 1}, []int{1, 1, 1, 1}, []int{1, 1, 1, 1}}, 8},
}

func TestEvalDiagonalSum(t *testing.T) {
	for _, pair := range diagonalSumTestCases {
		var newOut = diagonalSum(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
