package leetcode

import (
	"reflect"
	"testing"
)

type rangeSumBSTTestPair struct {
	inputRoot *TreeNode
	inputLow  int
	inputHigh int
	out       int
}

var rangeSumBSTTestCases = []rangeSumBSTTestPair{
	{generateTreeNode([]interface{}{10, 5, 15, 3, 7, nil, 18}), 7, 15, 32},
	{generateTreeNode([]interface{}{10, 5, 15, 3, 7, 13, 18, 1, nil, 6}), 6, 10, 23},
}

func TestEvalRangeSumBST(t *testing.T) {
	for _, pair := range rangeSumBSTTestCases {
		newOut := rangeSumBST(pair.inputRoot, pair.inputLow, pair.inputHigh)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputRoot,
				pair.inputLow, pair.inputHigh,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
