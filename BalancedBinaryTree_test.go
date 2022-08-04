package leetcode

import (
	"reflect"
	"testing"
)

type isBalancedTestPair struct {
	input *TreeNode
	out   bool
}

var isBalancedTestCases = []isBalancedTestPair{
	{generateTreeNode([]interface{}{3, 9, 20, nil, nil, 15, 7}), true},
	{generateTreeNode([]interface{}{1, 2, 2, 3, 3, nil, nil, 4, 4}), false},
	{generateTreeNode([]interface{}{}), true},
}

func TestEvalIsBalanced(t *testing.T) {
	for _, pair := range isBalancedTestCases {
		var newOut = isBalanced(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
