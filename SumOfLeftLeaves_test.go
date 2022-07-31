package leetcode

import (
	"reflect"
	"testing"
)

type sumOfLeftLeavesTestPair struct {
	input *TreeNode
	out   int
}

var sumOfLeftLeavesTestCases = []sumOfLeftLeavesTestPair{
	{generateTreeNode([]interface{}{3, 9, 20, nil, nil, 15, 7}), 24},
	{generateTreeNode([]interface{}{1}), 0},
}

func TestEvalSumOfLeftLeaves(t *testing.T) {
	for _, pair := range sumOfLeftLeavesTestCases {
		var newOut = sumOfLeftLeaves(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
