package leetcode

import (
	"reflect"
	"testing"
)

type middleNodeTestPair struct {
	input *ListNode
	out   *ListNode
}

var middleNodeTestCases = []middleNodeTestPair{
	{generateListNode([]int{1, 2, 3, 4, 5}), generateListNode([]int{3, 4, 5})},
	{generateListNode([]int{1, 2, 3, 4, 5, 6}), generateListNode([]int{4, 5, 6})},
}

func TestEvalMiddleNode(t *testing.T) {
	for _, pair := range middleNodeTestCases {
		var newOut = middleNode(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
