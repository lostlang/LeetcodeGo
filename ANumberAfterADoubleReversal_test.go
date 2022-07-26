package leetcode

import (
	"reflect"
	"testing"
)

type isSameAfterReversalsTestPair struct {
	input int
	out   bool
}

var isSameAfterReversalsTestCases = []isSameAfterReversalsTestPair{
	{526, true},
	{1800, false},
	{0, true},
}

func TestEvalIsSameAfterReversals(t *testing.T) {
	for _, pair := range isSameAfterReversalsTestCases {
		var newOut = isSameAfterReversals(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
