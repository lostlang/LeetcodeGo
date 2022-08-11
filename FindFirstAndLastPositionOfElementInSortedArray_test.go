package leetcode

import (
	"reflect"
	"testing"
)

type searchRangeTestPair struct {
	inputArr    []int
	inputTarget int
	out         []int
}

var searchRangeTestCases = []searchRangeTestPair{
	{[]int{}, 0, []int{-1, -1}},
	{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
	{[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
	{[]int{1, 1}, 1, []int{0, 1}},
}

func TestEvalSearchRange(t *testing.T) {
	for _, pair := range searchRangeTestCases {
		var newOut = searchRange(pair.inputArr, pair.inputTarget)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputArr,
				"and", pair.inputTarget,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
