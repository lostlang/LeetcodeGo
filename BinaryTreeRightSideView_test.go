package leetcode

import (
	"reflect"
	"testing"
)

type rightSideViewTestPair struct {
	input *TreeNode
	out   []int
}

var rightSideViewTestCases = []rightSideViewTestPair{
	{generateTreeNode([]interface{}{1, 2, 3, nil, 5, nil, 4}), []int{1, 3, 4}},
	{generateTreeNode([]interface{}{1, nil, 3}), []int{1, 3}},
	{generateTreeNode([]interface{}{}), []int{}},
	{generateTreeNode([]interface{}{1, 2, 3, 4}), []int{1, 3, 4}},
	{generateTreeNode([]interface{}{1, 2}), []int{1, 2}},
}

func TestEvalRightSideView(t *testing.T) {
	for _, pair := range rightSideViewTestCases {
		var newOut = rightSideView(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
