package leetcode

import (
	"reflect"
	"testing"
)

type searchRotateTestPair struct {
	inputArray  []int
	inputTarget int
	out         int
}

var searchRotateTestCases = []searchRotateTestPair{
	{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
	{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
	{[]int{1}, 0, -1},
}

func TestEvalSearchRotate(t *testing.T) {
	for _, pair := range searchRotateTestCases {
		var newOut = searchRotate(pair.inputArray, pair.inputTarget)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputArray, pair.inputTarget,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
