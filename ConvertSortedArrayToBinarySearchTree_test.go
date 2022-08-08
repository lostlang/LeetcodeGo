package leetcode

import (
	"reflect"
	"testing"
)

type sortedArrayToBSTTestPair struct {
	input []int
	out   *TreeNode
}

var sortedArrayToBSTTestCases = []sortedArrayToBSTTestPair{
	{[]int{-10, -3, 0, 5, 9}, generateTreeNode([]interface{}{0, -3, 9, -10, nil, 5})},
	{[]int{1, 3}, generateTreeNode([]interface{}{3, 1})},
}

func TestEvalSortedArrayToBST(t *testing.T) {
	for _, pair := range sortedArrayToBSTTestCases {
		var newOut = sortedArrayToBST(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
