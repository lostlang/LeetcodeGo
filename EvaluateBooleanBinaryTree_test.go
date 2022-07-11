package leetcode

import (
	"reflect"
	"testing"
)

type evaluateTreeTestPair struct {
	input *TreeNode
	out   bool
}

var evaluateTreeTestCases = []evaluateTreeTestPair{
	{generateTreeNode([]interface{}{2, 1, 3, nil, nil, 0, 1}), true},
	{generateTreeNode([]interface{}{0}), false},
}

func TestEvalEvaluateTree(t *testing.T) {
	for _, pair := range evaluateTreeTestCases {
		var newOut = evaluateTree(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
