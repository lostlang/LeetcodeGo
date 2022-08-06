package leetcode

import (
	"reflect"
	"testing"
)

type removeNthFromEndTestPair struct {
	inputList *ListNode
	inputN    int
	out       *ListNode
}

var removeNthFromEndTestCases = []removeNthFromEndTestPair{
	{generateListNode([]int{1, 2, 3, 4, 5}), 2, generateListNode([]int{1, 2, 3, 5})},
	{generateListNode([]int{1}), 1, nil},
	{generateListNode([]int{1, 2}), 1, generateListNode([]int{1})},
}

func TestEvalRemoveNthFromEnd(t *testing.T) {
	for _, pair := range removeNthFromEndTestCases {
		var newOut = removeNthFromEnd(pair.inputList, pair.inputN)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputList,
				"and", pair.inputN,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
