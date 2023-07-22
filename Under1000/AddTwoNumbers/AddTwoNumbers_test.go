package AddTwoNumbers

import (
	"reflect"
	"testing"

	"leetcode/utils"
)

type addTwoNumbersTestPair struct {
	inputL1 *ListNode
	inputL2 *ListNode
	output  *ListNode
}

var addTwoNumbersTestCases = []addTwoNumbersTestPair{
	{utils.NewListNode([]int{2, 4, 3}), utils.NewListNode([]int{5, 6, 4}), utils.NewListNode([]int{7, 0, 8})},
	{utils.NewListNode([]int{0}), utils.NewListNode([]int{0}), utils.NewListNode([]int{0})},
	{utils.NewListNode([]int{9, 9, 9, 9, 9, 9, 9}), utils.NewListNode([]int{9, 9, 9, 9}), utils.NewListNode([]int{8, 9, 9, 9, 0, 0, 0, 1})},
}

func TestEvalAddTwoNumbers(t *testing.T) {
	for _, pair := range addTwoNumbersTestCases {
		newOutput := addTwoNumbers(pair.inputL1, pair.inputL2)
		if !reflect.DeepEqual(newOutput, pair.output) {
			t.Error(
				"For", pair.inputL1, pair.inputL2,
				"expected", pair.output,
				"got", newOutput,
			)
		}
	}
}
