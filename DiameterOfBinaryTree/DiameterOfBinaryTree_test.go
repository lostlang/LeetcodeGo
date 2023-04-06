package DiameterOfBinaryTree

import (
	"reflect"
	"testing"

	"leetcode/utils"
)

type diameterOfBinaryTreeTestPair struct {
	input  *TreeNode
	output int
}

var diameterOfBinaryTreeTestCases = []diameterOfBinaryTreeTestPair{
	{utils.NewTreeNode([]interface{}{1, 2, 3, 4, 5}), 3},
	{utils.NewTreeNode([]interface{}{1, 2}), 1},
	{utils.NewTreeNode([]interface{}{4, -7, -3, nil, nil, -9, -3, 9, -7, -4, nil, 6, nil, -6, -6, nil, nil, 0, 6, 5, nil, 9, nil, nil, -1, -4, nil, nil, nil, -2}), 8},
}

func TestEvalDiameterOfBinaryTree(t *testing.T) {
	for _, pair := range diameterOfBinaryTreeTestCases {
		newOutput := diameterOfBinaryTree(pair.input)
		if !reflect.DeepEqual(newOutput, pair.output) {
			t.Error(
				"For", pair.input,
				"expected", pair.output,
				"got", newOutput,
			)
		}
	}
}
