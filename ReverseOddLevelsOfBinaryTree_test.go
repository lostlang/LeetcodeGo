package leetcode

import (
	"reflect"
	"testing"
)

type reverseOddLevelsTestPair struct {
	input *TreeNode
	out   *TreeNode
}

var reverseOddLevelsTestCases = []reverseOddLevelsTestPair{
	{generateTreeNode([]interface{}{2, 3, 5, 8, 13, 21, 34}), generateTreeNode([]interface{}{2, 5, 3, 8, 13, 21, 34})},
	{generateTreeNode([]interface{}{7, 13, 11}), generateTreeNode([]interface{}{7, 11, 13})},
	{generateTreeNode([]interface{}{0, 1, 2, 0, 0, 0, 0, 1, 1, 1, 1, 2, 2, 2, 2}), generateTreeNode([]interface{}{0, 2, 1, 0, 0, 0, 0, 2, 2, 2, 2, 1, 1, 1, 1})},
}

func TestEvalReverseOddLevels(t *testing.T) {
	for _, pair := range reverseOddLevelsTestCases {
		var newOut = reverseOddLevels(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
