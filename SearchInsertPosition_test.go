package leetcode

import (
	"reflect"
	"testing"
)

type searchInsertTestPair struct {
	inputArray  []int
	inputTarget int
	out         int
}

var searchInsertTestCases = []searchInsertTestPair{
	{[]int{1, 3, 5, 6}, 5, 2},
	{[]int{1, 3, 5, 6}, 2, 1},
	{[]int{1, 3, 5, 6}, 7, 4},
}

func TestEvalSearchInsert(t *testing.T) {
	for _, pair := range searchInsertTestCases {
		var newOut = searchInsert(pair.inputArray, pair.inputTarget)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputArray,
				"and", pair.inputTarget,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
