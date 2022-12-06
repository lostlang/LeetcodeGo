package leetcode

import (
	"reflect"
	"testing"
)

type minPathSumTestPair struct {
	input [][]int
	out   int
}

var minPathSumTestCases = []minPathSumTestPair{
	{[][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}, 7},
	{[][]int{{1, 2, 3}, {4, 5, 6}}, 12},
}

func TestEvalMinPathSum(t *testing.T) {
	for _, pair := range minPathSumTestCases {
		var newOut = minPathSum(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
