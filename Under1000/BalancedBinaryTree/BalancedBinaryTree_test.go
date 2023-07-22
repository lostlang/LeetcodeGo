package BalancedBinaryTree

import (
	"reflect"
	"testing"

	"leetcode/utils"
)

type isBalancedTestPair struct {
	input  *TreeNode
	output bool
}

var isBalancedTestCases = []isBalancedTestPair{
	{utils.NewTreeNode([]interface{}{3, 9, 20, nil, nil, 15, 7}), true},
	{utils.NewTreeNode([]interface{}{1, 2, 2, 3, 3, nil, nil, 4, 4}), false},
	{utils.NewTreeNode([]interface{}{}), true},
}

func TestEvalIsBalanced(t *testing.T) {
	for _, pair := range isBalancedTestCases {
		newOutput := isBalanced(pair.input)
		if !reflect.DeepEqual(newOutput, pair.output) {
			t.Error(
				"For", pair.input,
				"expected", pair.output,
				"got", newOutput,
			)
		}
	}
}
