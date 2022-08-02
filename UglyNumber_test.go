package leetcode

import (
	"reflect"
	"testing"
)

type isUglyTestPair struct {
	input int
	out   bool
}

var isUglyTestCases = []isUglyTestPair{
	{6, true},
	{1, true},
	{14, false},
}

func TestEvalIsUgly(t *testing.T) {
	for _, pair := range isUglyTestCases {
		var newOut = isUgly(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
