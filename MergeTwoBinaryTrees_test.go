package leetcode

import (
	"reflect"
	"testing"
)

type mergeTreesTestPair struct {
	inputRoot1 *TreeNode
	inputRoot2 *TreeNode
	out        *TreeNode
}

var mergeTreesTestCases = []mergeTreesTestPair{
	{generateTreeNode([]interface{}{1, 3, 2, 5}), generateTreeNode([]interface{}{2, 1, 3, nil, 4, nil, 7}), generateTreeNode([]interface{}{3, 4, 5, 5, 4, nil, 7})},
	{generateTreeNode([]interface{}{1}), generateTreeNode([]interface{}{1, 2}), generateTreeNode([]interface{}{2, 2})},
}

func TestEvalMergeTrees(t *testing.T) {
	for _, pair := range mergeTreesTestCases {
		var newOut = mergeTrees(pair.inputRoot1, pair.inputRoot2)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputRoot1,
				"and", pair.inputRoot2,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
