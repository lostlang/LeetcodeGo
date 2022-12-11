package leetcode

import (
	"reflect"
	"testing"
)

type getMaxLenTestPair struct {
	input []int
	out   int
}

var getMaxLenTestCases = []getMaxLenTestPair{
	{[]int{1, -2, -3, 4}, 4},
	{[]int{0, 1, -2, -3, -4}, 3},
	{[]int{-1, -2, -3, 0, 1}, 2},
}

func TestEvalGetMaxLen(t *testing.T) {
	for _, pair := range getMaxLenTestCases {
		var newOut = getMaxLen(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
