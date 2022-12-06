package leetcode

import (
	"reflect"
	"testing"
)

type addTwoNumbersTestPair struct {
	inputL1 *ListNode
	inputL2 *ListNode
	out     *ListNode
}

var addTwoNumbersTestCases = []addTwoNumbersTestPair{
	{generateListNode([]int{2, 4, 3}), generateListNode([]int{5, 6, 4}), generateListNode([]int{7, 0, 8})},
	{generateListNode([]int{0}), generateListNode([]int{0}), generateListNode([]int{0})},
	{generateListNode([]int{9, 9, 9, 9, 9, 9, 9}), generateListNode([]int{9, 9, 9, 9}), generateListNode([]int{8, 9, 9, 9, 0, 0, 0, 1})},
}

func TestEvalAddTwoNumbers(t *testing.T) {
	for _, pair := range addTwoNumbersTestCases {
		var newOut = addTwoNumbers(pair.inputL1, pair.inputL2)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputL1, pair.inputL2,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
