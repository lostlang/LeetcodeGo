package leetcode

import (
	"reflect"
	"testing"
)

type hasPathSumTestPair struct {
	inputTree   *TreeNode
	inputTargen int
	out         bool
}

var hasPathSumTestCases = []hasPathSumTestPair{
	{generateTreeNode([]interface{}{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, nil, 1}), 22, true},
	{generateTreeNode([]interface{}{1, 2, 3}), 5, false},
	{generateTreeNode([]interface{}{}), 0, false},
}

func TestEvalHasPathSum(t *testing.T) {
	for _, pair := range hasPathSumTestCases {
		var newOut = hasPathSum(pair.inputTree, pair.inputTargen)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputTree,
				"and", pair.inputTargen,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
