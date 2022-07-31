package leetcode

import (
	"reflect"
	"testing"
)

type searchBSTTestPair struct {
	inputTree   *TreeNode
	inputTargen int
	out         *TreeNode
}

var searchBSTTestCases = []searchBSTTestPair{
	{generateTreeNode([]interface{}{4, 2, 7, 1, 3}), 5, generateTreeNode([]interface{}{})},
	{generateTreeNode([]interface{}{4, 2, 7, 1, 3}), 2, generateTreeNode([]interface{}{2, 1, 3})},
}

func TestEvalSearchBST(t *testing.T) {
	for _, pair := range searchBSTTestCases {
		var newOut = searchBST(pair.inputTree, pair.inputTargen)
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
