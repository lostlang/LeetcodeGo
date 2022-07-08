package leetcode

import (
	"reflect"
	"testing"
)

type deleteNodeTestPair struct {
	input *ListNode
	out   *ListNode
}

var deleteNodeTestCases = []deleteNodeTestPair{
	{generateListNode([]int{5, 1, 9}), generateListNode([]int{1, 9})},
	{generateListNode([]int{1, 9}), generateListNode([]int{9})},
}

func TestEvalDeleteNode(t *testing.T) {
	for _, pair := range deleteNodeTestCases {
		deleteNode(pair.input)
		if !reflect.DeepEqual(pair.input, pair.out) {
			t.Error(
				"Got", pair.input,
				"expected", pair.out)
		}
	}
}
