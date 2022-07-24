package leetcode

import (
	"reflect"
	"testing"
)

type preorderTestPair struct {
	input *Node
	out   []int
}

var preorderTestCases = []preorderTestPair{
	{generateNode([]interface{}{1, nil, 3, 2, 4, nil, 5, 6}), []int{1, 3, 5, 6, 2, 4}},
	{generateNode([]interface{}{1, nil, 2, 3, 4, 5, nil, nil, 6, 7, nil, 8, nil, 9, 10, nil, nil, 11, nil, 12, nil, 13, nil, nil, 14}), []int{1, 2, 3, 6, 7, 11, 14, 4, 8, 12, 5, 9, 13, 10}},
}

func TestEvalPreorder(t *testing.T) {
	for _, pair := range preorderTestCases {
		var newOut = preorder(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
