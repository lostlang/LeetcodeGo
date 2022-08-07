package leetcode

import (
	"reflect"
	"testing"
)

type minDepthTestPair struct {
	input *TreeNode
	out   int
}

var minDepthTestCases = []minDepthTestPair{
	{generateTreeNode([]interface{}{3, 9, 20, nil, nil, 15, 7}), 2},
	{generateTreeNode([]interface{}{2, nil, 3, nil, 4, nil, 5, nil, 6}), 5},
}

func TestEvalMinDepth(t *testing.T) {
	for _, pair := range minDepthTestCases {
		var newOut = minDepth(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
