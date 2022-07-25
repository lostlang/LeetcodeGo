package leetcode

import (
	"reflect"
	"testing"
)

type arraySignTestPair struct {
	input []int
	out   int
}

var arraySignTestCases = []arraySignTestPair{
	{[]int{-1, -2, -3, -4, 3, 2, 1}, 1},
	{[]int{1, 5, 0, 2, -3}, 0},
	{[]int{-1, 1, -1, 1, -1}, -1},
}

func TestEvalArraySign(t *testing.T) {
	for _, pair := range arraySignTestCases {
		var newOut = arraySign(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
