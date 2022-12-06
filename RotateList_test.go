package leetcode

import (
	"reflect"
	"testing"
)

type rotateRightTestPair struct {
	inputHead *ListNode
	inputK    int
	out       *ListNode
}

var rotateRightTestCases = []rotateRightTestPair{
	{generateListNode([]int{1, 2, 3, 4, 5}), 2, generateListNode([]int{4, 5, 1, 2, 3})},
	{generateListNode([]int{0, 1, 2}), 4, generateListNode([]int{2, 0, 1})},
	{generateListNode([]int{1, 2}), 0, generateListNode([]int{1, 2})},
	{generateListNode([]int{1, 2}), 1, generateListNode([]int{2, 1})},
}

func TestEvalRotateRight(t *testing.T) {
	for _, pair := range rotateRightTestCases {
		var newOut = rotateRight(pair.inputHead, pair.inputK)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputHead, pair.inputK,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
