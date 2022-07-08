package leetcode

import (
	"reflect"
	"testing"
)

type removeElementsTestPair struct {
	inputList   *ListNode
	inputDelete int
	out         *ListNode
}

var removeElementsTestCases = []removeElementsTestPair{
	{generateListNode([]int{1, 2, 6, 3, 4, 5, 6}), 6, generateListNode([]int{1, 2, 3, 4, 5})},
	{generateListNode([]int{}), 1, generateListNode([]int{})},
	{generateListNode([]int{7, 7, 7, 7}), 7, generateListNode([]int{})},
}

func TestEvalRemoveElements(t *testing.T) {
	for _, pair := range removeElementsTestCases {
		var newOut = removeElements(pair.inputList, pair.inputDelete)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputList,
				"and", pair.inputDelete,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
