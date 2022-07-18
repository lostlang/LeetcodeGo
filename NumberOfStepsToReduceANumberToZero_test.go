package leetcode

import (
	"reflect"
	"testing"
)

type numberOfStepsTestPair struct {
	input int
	out   int
}

var numberOfStepsTestCases = []numberOfStepsTestPair{
	{14, 6},
	{8, 4},
	{123, 12},
}

func TestEvalNumberOfSteps(t *testing.T) {
	for _, pair := range numberOfStepsTestCases {
		var newOut = numberOfSteps(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
