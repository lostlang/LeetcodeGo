package BinarySearch

import (
	"reflect"
	"testing"
)

type searchTestPair struct {
	inputArray  []int
	inputTarget int
	output      int
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
		newOutput := search(pair.inputArray, pair.inputTarget)
		if !reflect.DeepEqual(newOutput, pair.output) {
			t.Error(
				"For", pair.inputArray, pair.inputTarget,
				"expected", pair.output,
				"got", newOutput,
			)
		}
	}
}
