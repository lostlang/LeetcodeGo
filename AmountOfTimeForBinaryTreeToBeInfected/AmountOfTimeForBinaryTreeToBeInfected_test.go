package leetcode

import (
	"reflect"
	"testing"
)

type amountOfTimeTestPair struct {
	inputRoot  *TreeNode
	inputStart int
	output     int
}

var amountOfTimeTestCases = []amountOfTimeTestPair{
	// {generateTreeNode([]interface{}{1, 5, 3, nil, 4, 10, 6, 9, 2}), 3, 4},
	// {generateTreeNode([]interface{}{1}), 1, 0},
	// {generateTreeNode([]interface{}{1, 2, nil, 3, nil, 4, nil, 5}), 3, 2},
	// {generateTreeNode([]interface{}{5, 2, 3, 4, nil, nil, nil, 1}), 4, 3},
}

func TestEvalAmountOfTime(t *testing.T) {
	for _, pair := range amountOfTimeTestCases {
		newOutput := amountOfTime(pair.inputRoot, pair.inputStart)
		if !reflect.DeepEqual(newOutput, pair.output) {
			t.Error(
				"For", pair.inputRoot, pair.inputStart,
				"expected", pair.output,
				"got", newOutput,
			)
		}
	}
}
