package leetcode

import (
	"reflect"
	"testing"
)

type isHappyTestPair struct {
	input int
	out   bool
}

var isHappyTestCases = []isHappyTestPair{
	{19, true},
	{2, false},
}

func TestEvalIsHappy(t *testing.T) {
	for _, pair := range isHappyTestCases {
		var newOut = isHappy(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
