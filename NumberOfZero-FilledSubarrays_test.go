package leetcode

import (
	"reflect"
	"testing"
)

type zeroFilledSubarrayTestPair struct {
	input []int
	out   int64
}

var zeroFilledSubarrayTestCases = []zeroFilledSubarrayTestPair{
	{[]int{1, 3, 0, 0, 2, 0, 0, 4}, 6},
	{[]int{0, 0, 0, 2, 0, 0}, 9},
	{[]int{2, 10, 2019}, 0},
}

func TestEvalZeroFilledSubarray(t *testing.T) {
	for _, pair := range zeroFilledSubarrayTestCases {
		var newOut = zeroFilledSubarray(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
