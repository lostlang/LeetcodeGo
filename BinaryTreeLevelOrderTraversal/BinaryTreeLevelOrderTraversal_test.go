package BinaryTreeLevelOrderTraversal

import (
	"reflect"
	"testing"

	"leetcode/utils"
)

type levelOrderTestPair struct {
	input  *TreeNode
	output [][]int
}

var levelOrderTestCases = []levelOrderTestPair{
	{utils.NewTreeNode([]interface{}{3, 9, 20, nil, nil, 15, 7}), [][]int{{3}, {9, 20}, {15, 7}}},
	{utils.NewTreeNode([]interface{}{}), [][]int{}},
	{utils.NewTreeNode([]interface{}{1}), [][]int{{1}}},
}

func TestEvalLevelOrder(t *testing.T) {
	for _, pair := range levelOrderTestCases {
		newOutput := levelOrder(pair.input)
		if !reflect.DeepEqual(newOutput, pair.output) {
			t.Error(
				"For", pair.input,
				"expected", pair.output,
				"got", newOutput,
			)
		}
	}
}
