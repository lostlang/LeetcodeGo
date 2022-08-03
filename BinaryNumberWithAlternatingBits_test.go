package leetcode

import (
	"reflect"
	"testing"
)

type hasAlternatingBitsTestPair struct {
	input int
	out   bool
}

var hasAlternatingBitsTestCases = []hasAlternatingBitsTestPair{
	{5, true},
	{7, false},
	{11, false},
}

func TestEvalHasAlternatingBits(t *testing.T) {
	for _, pair := range hasAlternatingBitsTestCases {
		var newOut = hasAlternatingBits(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
