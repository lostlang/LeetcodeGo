package leetcode

import (
	"reflect"
	"testing"
)

type maxDepthTestPair struct {
	input *TreeNode
	out   int
}

var maxDepthTestCases = []maxDepthTestPair{
	{generateTreeNode([]interface{}{3, 9, 20, nil, nil, 15, 7}), 3},
	{generateTreeNode([]interface{}{1, nil, 2}), 2},
}

func TestEvalMaxDepth(t *testing.T) {
	for _, pair := range maxDepthTestCases {
		var newOut = maxDepth(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
