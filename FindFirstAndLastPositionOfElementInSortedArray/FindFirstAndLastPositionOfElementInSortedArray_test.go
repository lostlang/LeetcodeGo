package FindFirstAndLastPositionOfElementInSortedArray

import (
	"reflect"
	"testing"
)

type searchRangeTestPair struct {
	inputArr    []int
	inputTarget int
	output      []int
}

var searchRangeTestCases = []searchRangeTestPair{
	{[]int{}, 0, []int{-1, -1}},
	{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
	{[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
	{[]int{1, 1}, 1, []int{0, 1}},
}

func TestEvalSearchRange(t *testing.T) {
	for _, pair := range searchRangeTestCases {
		newOutput := searchRange(pair.inputArr, pair.inputTarget)
		if !reflect.DeepEqual(newOutput, pair.output) {
			t.Error(
				"For", pair.inputArr, pair.inputTarget,
				"expected", pair.output,
				"got", newOutput,
			)
		}
	}
}
