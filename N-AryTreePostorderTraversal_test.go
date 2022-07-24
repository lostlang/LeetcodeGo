package leetcode

import (
	"reflect"
	"testing"
)

type postorderTestPair struct {
	input *Node
	out   []int
}

var postorderTestCases = []postorderTestPair{
	{generateNode([]interface{}{1, nil, 3, 2, 4, nil, 5, 6}), []int{5, 6, 3, 2, 4, 1}},
	{generateNode([]interface{}{1, nil, 2, 3, 4, 5, nil, nil, 6, 7, nil, 8, nil, 9, 10, nil, nil, 11, nil, 12, nil, 13, nil, nil, 14}), []int{2, 6, 14, 11, 7, 3, 12, 8, 4, 13, 9, 10, 5, 1}},
}

func TestEvalPostorder(t *testing.T) {
	for _, pair := range postorderTestCases {
		var newOut = postorder(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
