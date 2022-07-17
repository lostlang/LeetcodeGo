package leetcode

import (
	"reflect"
	"testing"
)

type minimumSumTestPair struct {
	input int
	out   int
}

var minimumSumTestCases = []minimumSumTestPair{
	{2932, 52},
	{4009, 13},
}

func TestEvalMinimumSum(t *testing.T) {
	for _, pair := range minimumSumTestCases {
		var newOut = minimumSum(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
