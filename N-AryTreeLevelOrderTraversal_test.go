package leetcode

import (
	"reflect"
	"testing"
)

type levelOrderNodesTestPair struct {
	input *Node
	out   [][]int
}

var levelOrderNodesTestCases = []levelOrderNodesTestPair{
	{generateNode([]interface{}{1, nil, 3, 2, 4, nil, 5, 6}), [][]int{[]int{1}, []int{3, 2, 4}, []int{5, 6}}},
}

func TestEvalLevelOrderNodes(t *testing.T) {
	for _, pair := range levelOrderNodesTestCases {
		var newOut = levelOrderNodes(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
