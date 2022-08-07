package leetcode

import (
	"reflect"
	"testing"
)

type levelOrderTestPair struct {
	input *TreeNode
	out   [][]int
}

var levelOrderTestCases = []levelOrderTestPair{
	{generateTreeNode([]interface{}{3, 9, 20, nil, nil, 15, 7}), [][]int{{3}, {9, 20}, {15, 7}}},
	{generateTreeNode([]interface{}{}), [][]int{}},
	{generateTreeNode([]interface{}{1}), [][]int{{1}}},
}

func TestEvalLevelOrder(t *testing.T) {
	for _, pair := range levelOrderTestCases {
		var newOut = levelOrder(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
