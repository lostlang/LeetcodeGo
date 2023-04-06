package MaximumDepthOfBinaryTree

import (
	"reflect"
	"testing"

	"leetcode/utils"
)

type maxDepthTestPair struct {
	input  *TreeNode
	output int
}

var maxDepthTestCases = []maxDepthTestPair{
	{utils.NewTreeNode([]interface{}{3, 9, 20, nil, nil, 15, 7}), 3},
	{utils.NewTreeNode([]interface{}{1, nil, 2}), 2},
}

func TestEvalMaxDepth(t *testing.T) {
	for _, pair := range maxDepthTestCases {
		newOutput := maxDepth(pair.input)
		if !reflect.DeepEqual(newOutput, pair.output) {
			t.Error(
				"For", pair.input,
				"expected", pair.output,
				"got", newOutput,
			)
		}
	}
}
