package leetcode

import (
	"reflect"
	"testing"
)

type searchTestPair struct {
	inputArray  []int
	inputTarget int
	out         int
}

var searchTestCases = []searchTestPair{
	{[]int{-1, 0, 3, 5, 9, 12}, 9, 4},
	{[]int{-1, 0, 3, 5, 9, 12}, 2, -1},
	{[]int{2, 5}, 2, 0},
	{[]int{-5}, 5, -1},
	{[]int{-1, 0, 5}, 2, -1},
}

func TestEvalSearch(t *testing.T) {
	for _, pair := range searchTestCases {
		var newOut = search(pair.inputArray, pair.inputTarget)
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
