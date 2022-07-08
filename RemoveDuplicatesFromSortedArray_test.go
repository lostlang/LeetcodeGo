package leetcode

import (
	"reflect"
	"testing"
)

type removeDuplicatesTestPair struct {
	input []int
	out   int
}

var removeDuplicatesTestCases = []removeDuplicatesTestPair{
	{[]int{1, 1, 2}, 2},
	{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
}

func TestEvalRemoveDuplicates(t *testing.T) {
	for _, pair := range removeDuplicatesTestCases {
		var newOut = removeDuplicates(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
