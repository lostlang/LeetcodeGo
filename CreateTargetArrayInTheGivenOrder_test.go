package leetcode

import (
	"reflect"
	"testing"
)

type createTargetArrayTestPair struct {
	inputNums  []int
	inputIndex []int
	out        []int
}

var createTargetArrayTestCases = []createTargetArrayTestPair{
	{[]int{0, 1, 2, 3, 4}, []int{0, 1, 2, 2, 1}, []int{0, 4, 1, 3, 2}},
	{[]int{1, 2, 3, 4, 0}, []int{0, 1, 2, 3, 0}, []int{0, 1, 2, 3, 4}},
	{[]int{1}, []int{0}, []int{1}},
}

func TestEvalCreateTargetArray(t *testing.T) {
	for _, pair := range createTargetArrayTestCases {
		var newOut = createTargetArray(pair.inputNums, pair.inputIndex)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputNums,
				"and", pair.inputIndex,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
