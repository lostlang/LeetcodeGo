package leetcode

import (
	"reflect"
	"testing"
)

type isSymmetricTestPair struct {
	input *TreeNode
	out   bool
}

var isSymmetricTestCases = []isSymmetricTestPair{
	{generateTreeNode([]interface{}{1, 2, 2, 3, 4, 4, 3}), true},
	{generateTreeNode([]interface{}{1, 2, 2, nil, 3, nil, 3}), false},
	{generateTreeNode([]interface{}{1, 2}), false},
	{generateTreeNode([]interface{}{1, 2, 3}), false},
	{generateTreeNode([]interface{}{1}), true},
}

func TestEvalIsSymmetric(t *testing.T) {
	for _, pair := range isSymmetricTestCases {
		var newOut = isSymmetric(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
