package leetcode

import (
	"reflect"
	"testing"
)

type diameterOfBinaryTreeTestPair struct {
	input *TreeNode
	out   int
}

var diameterOfBinaryTreeTestCases = []diameterOfBinaryTreeTestPair{
	{generateTreeNode([]interface{}{1, 2, 3, 4, 5}), 3},
	{generateTreeNode([]interface{}{1, 2}), 1},
	{generateTreeNode([]interface{}{4, -7, -3, nil, nil, -9, -3, 9, -7, -4, nil, 6, nil, -6, -6, nil, nil, 0, 6, 5, nil, 9, nil, nil, -1, -4, nil, nil, nil, -2}), 7},
}

func TestEvalDiameterOfBinaryTree(t *testing.T) {
	for _, pair := range diameterOfBinaryTreeTestCases {
		var newOut = diameterOfBinaryTree(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
