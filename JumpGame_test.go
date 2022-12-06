package leetcode

import (
	"reflect"
	"testing"
)

type canJumpTestPair struct {
	input []int
	out   bool
}

var canJumpTestCases = []canJumpTestPair{
	{[]int{2, 3, 1, 1, 4}, true},
	{[]int{3, 2, 1, 0, 4}, false},
	{[]int{0}, true},
}

func TestEvalCanJump(t *testing.T) {
	for _, pair := range canJumpTestCases {
		var newOut = canJump(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
