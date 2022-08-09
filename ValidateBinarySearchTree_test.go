package leetcode

import (
	"reflect"
	"testing"
)

type isValidBSTTestPair struct {
	input *TreeNode
	out   bool
}

var isValidBSTTestCases = []isValidBSTTestPair{
	{generateTreeNode([]interface{}{2, 1, 3}), true},
	{generateTreeNode([]interface{}{5, 1, 4, nil, nil, 3, 6}), false},
}

func TestEvalIsValidBST(t *testing.T) {
	for _, pair := range isValidBSTTestCases {
		var newOut = isValidBST(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
