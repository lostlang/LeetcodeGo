package leetcode

import (
	"reflect"
	"testing"
)

type checkTreeTestPair struct {
	input *TreeNode
	out   bool
}

var checkTreeTestCases = []checkTreeTestPair{
	{generateTreeNode([]interface{}{10, 4, 6}), true},
	{generateTreeNode([]interface{}{5, 3, 1}), false},
}

func TestEvalCheckTree(t *testing.T) {
	for _, pair := range checkTreeTestCases {
		var newOut = checkTree(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
