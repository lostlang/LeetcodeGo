package leetcode

import (
	"reflect"
	"testing"
)

type binaryTreePathsTestPair struct {
	input *TreeNode
	out   []string
}

var binaryTreePathsTestCases = []binaryTreePathsTestPair{
	{generateTreeNode([]interface{}{1, 2, 3, nil, 5}), []string{"1->2->5", "1->3"}},
	{generateTreeNode([]interface{}{1}), []string{"1"}},
}

func TestEvalBinaryTreePaths(t *testing.T) {
	for _, pair := range binaryTreePathsTestCases {
		var newOut = binaryTreePaths(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
