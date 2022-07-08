package leetcode

import (
	"reflect"
	"testing"
)

type removeElementTestPair struct {
	inputArray  []int
	inputDelete int
	out         int
}

var removeElementTestCases = []removeElementTestPair{
	{[]int{3, 2, 2, 3}, 3, 2},
	{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5},
}

func TestEvalRemoveElement(t *testing.T) {
	for _, pair := range removeElementTestCases {
		var newOut = removeElement(pair.inputArray, pair.inputDelete)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputArray,
				"and", pair.inputDelete,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
