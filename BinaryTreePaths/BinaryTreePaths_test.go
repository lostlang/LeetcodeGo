package BinaryTreePaths

import (
	"reflect"
	"testing"

	"leetcode/utils"
)

type binaryTreePathsTestPair struct {
	input  *TreeNode
	output []string
}

var binaryTreePathsTestCases = []binaryTreePathsTestPair{
	{utils.NewTreeNode([]interface{}{1, 2, 3, nil, 5}), []string{"1->2->5", "1->3"}},
	{utils.NewTreeNode([]interface{}{1}), []string{"1"}},
}

func TestEvalBinaryTreePaths(t *testing.T) {
	for _, pair := range binaryTreePathsTestCases {
		newOutput := binaryTreePaths(pair.input)
		if !reflect.DeepEqual(newOutput, pair.output) {
			t.Error(
				"For", pair.input,
				"expected", pair.output,
				"got", newOutput,
			)
		}
	}
}
