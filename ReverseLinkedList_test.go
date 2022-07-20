package leetcode

import (
	"reflect"
	"testing"
)

type reverseListTestPair struct {
	input *ListNode
	out   *ListNode
}

var reverseListTestCases = []reverseListTestPair{
	{generateListNode([]int{1, 2, 3, 4, 5}), generateListNode([]int{5, 4, 3, 2, 1})},
	{generateListNode([]int{1, 2}), generateListNode([]int{2, 1})},
	{generateListNode([]int{}), generateListNode([]int{})},
}

func TestEvalReverseList(t *testing.T) {
	for _, pair := range reverseListTestCases {
		var newOut = reverseList(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
