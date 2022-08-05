package leetcode

import (
	"reflect"
	"testing"
)

type balancedStringSplitTestPair struct {
	input string
	out   int
}

var balancedStringSplitTestCases = []balancedStringSplitTestPair{
	{"RLRRLLRLRL", 4},
	{"RLRRRLLRLL", 2},
	{"LLLLRRRR", 1},
}

func TestEvalBalancedStringSplit(t *testing.T) {
	for _, pair := range balancedStringSplitTestCases {
		var newOut = balancedStringSplit(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
