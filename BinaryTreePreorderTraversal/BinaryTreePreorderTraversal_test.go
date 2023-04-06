package BinaryTreePreorderTraversal

import (
	"reflect"
	"testing"

	"leetcode/utils"
)

type preorderTraversalTestPair struct {
	input  *TreeNode
	output []int
}

var preorderTraversalTestCases = []preorderTraversalTestPair{
	{utils.NewTreeNode([]interface{}{1, nil, 2, 3}), []int{1, 2, 3}},
	{utils.NewTreeNode([]interface{}{}), []int{}},
	{utils.NewTreeNode([]interface{}{1}), []int{1}},
}

func TestEvalPreorderTraversal(t *testing.T) {
	for _, pair := range preorderTraversalTestCases {
		newOutput := preorderTraversal(pair.input)
		if !reflect.DeepEqual(newOutput, pair.output) {
			t.Error(
				"For", pair.input,
				"expected", pair.output,
				"got", newOutput,
			)
		}
	}
}
