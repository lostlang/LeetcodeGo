package leetcode

import (
	"reflect"
	"testing"
)

type getDecimalValueTestPair struct {
	input *ListNode
	out   int
}

var getDecimalValueTestCases = []getDecimalValueTestPair{
	{generateListNode([]int{1, 0, 1}), 5},
	{generateListNode([]int{0}), 0},
}

func TestEvalGetDecimalValue(t *testing.T) {
	for _, pair := range getDecimalValueTestCases {
		var newOut = getDecimalValue(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
