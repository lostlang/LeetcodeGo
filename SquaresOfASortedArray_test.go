package leetcode

import (
	"reflect"
	"testing"
)

type sortedSquaresTestPair struct {
	input []int
	out   []int
}

var sortedSquaresTestCases = []sortedSquaresTestPair{
	{[]int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100}},
	{[]int{-7, -3, 2, 3, 11}, []int{4, 9, 9, 49, 121}},
}

func TestEvalSortedSquares(t *testing.T) {
	for _, pair := range sortedSquaresTestCases {
		var newOut = sortedSquares(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
