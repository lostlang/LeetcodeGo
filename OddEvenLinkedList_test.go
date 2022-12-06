package leetcode

import (
	"reflect"
	"testing"
)

type oddEvenListTestPair struct {
	input *ListNode
	out   *ListNode
}

var oddEvenListTestCases = []oddEvenListTestPair{
	{generateListNode([]int{1, 2, 3, 4, 5}), generateListNode([]int{1, 3, 5, 2, 4})},
	{generateListNode([]int{2, 1, 3, 5, 6, 4, 7}), generateListNode([]int{2, 3, 6, 7, 1, 5, 4})},
}

func TestEvalOddEvenList(t *testing.T) {
	for _, pair := range oddEvenListTestCases {
		var newOut = oddEvenList(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
