package BinaryTreeRightSideView

import (
	"reflect"
	"testing"

	"leetcode/utils"
)

type rightSideViewTestPair struct {
	input *TreeNode
	out   []int
}

var rightSideViewTestCases = []rightSideViewTestPair{
	{utils.NewTreeNode([]interface{}{1, 2, 3, nil, 5, nil, 4}), []int{1, 3, 4}},
	{utils.NewTreeNode([]interface{}{1, nil, 3}), []int{1, 3}},
	{utils.NewTreeNode([]interface{}{}), []int{}},
	{utils.NewTreeNode([]interface{}{1, 2, 3, 4}), []int{1, 3, 4}},
	{utils.NewTreeNode([]interface{}{1, 2}), []int{1, 2}},
}

func TestEvalRightSideView(t *testing.T) {
	for _, pair := range rightSideViewTestCases {
		newOut := rightSideView(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
