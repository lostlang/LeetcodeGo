package leetcode

import (
	"reflect"
	"testing"
)

type mergeTwoListsTestPair struct {
	input_1 *ListNode
	input_2 *ListNode
	out     *ListNode
}

var mergeTwoListsTestCases = []mergeTwoListsTestPair{
	{generateListNode([]int{1, 2, 4}), generateListNode([]int{1, 3, 4}), generateListNode([]int{1, 1, 2, 3, 4, 4})},
	{generateListNode([]int{}), generateListNode([]int{}), generateListNode([]int{})},
	{generateListNode([]int{}), generateListNode([]int{0}), generateListNode([]int{0})},
}

func TestEvalMergeTwoLists(t *testing.T) {
	for _, pair := range mergeTwoListsTestCases {
		var newOut = mergeTwoLists(pair.input_1, pair.input_2)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input_1,
				"and", pair.input_2,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
