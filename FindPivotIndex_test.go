package leetcode

import (
	"reflect"
	"testing"
)

type pivotIndexTestPair struct {
	input []int
	out   int
}

var pivotIndexTestCases = []pivotIndexTestPair{
	{[]int{1, 7, 3, 6, 5, 6}, 3},
	{[]int{1, 2, 3}, -1},
	{[]int{2, 1, -1}, 0},
}

func TestEvalPivotIndex(t *testing.T) {
	for _, pair := range pivotIndexTestCases {
		var newOut = pivotIndex(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
