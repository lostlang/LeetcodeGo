package leetcode

import (
	"reflect"
	"testing"
)

type invertTreeTestPair struct {
	input *TreeNode
	out   *TreeNode
}

var invertTreeTestCases = []invertTreeTestPair{
	{generateTreeNode([]interface{}{4, 2, 7, 1, 3, 6, 9}), generateTreeNode([]interface{}{4, 7, 2, 9, 6, 3, 1})},
	{generateTreeNode([]interface{}{2, 1, 3}), generateTreeNode([]interface{}{2, 3, 1})},
	{generateTreeNode([]interface{}{}), generateTreeNode([]interface{}{})},
}

func TestEvalInvertTree(t *testing.T) {
	for _, pair := range invertTreeTestCases {
		var newOut = invertTree(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
