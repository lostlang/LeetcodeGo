package leetcode

import (
	"reflect"
	"testing"
)

type maxSubArrayTestPair struct {
	input []int
	out   int
}

var maxSubArrayTestCases = []maxSubArrayTestPair{
	{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
	{[]int{1}, 1},
	{[]int{5, 4, -1, 7, 8}, 23},
}

func TestEvalMaxSubArray(t *testing.T) {
	for _, pair := range maxSubArrayTestCases {
		var newOut = maxSubArray(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
