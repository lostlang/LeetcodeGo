package leetcode

import (
	"reflect"
	"testing"
)

type isPowerOfTwoTestPair struct {
	input int
	out   bool
}

var isPowerOfTwoTestCases = []isPowerOfTwoTestPair{
	{1, true},
	{16, true},
	{3, false},
}

func TestEvalIsPowerOfTwo(t *testing.T) {
	for _, pair := range isPowerOfTwoTestCases {
		var newOut = isPowerOfTwo(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
