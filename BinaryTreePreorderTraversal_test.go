package leetcode

import (
	"reflect"
	"testing"
)

type preorderTraversalTestPair struct {
	input *TreeNode
	out   []int
}

var preorderTraversalTestCases = []preorderTraversalTestPair{
	{generateTreeNode([]interface{}{1, nil, 2, 3}), []int{1, 2, 3}},
	{generateTreeNode([]interface{}{}), []int{}},
	{generateTreeNode([]interface{}{1}), []int{1}},
}

func TestEvalPreorderTraversal(t *testing.T) {
	for _, pair := range preorderTraversalTestCases {
		var newOut = preorderTraversal(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
