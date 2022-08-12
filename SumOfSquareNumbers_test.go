package leetcode

import (
	"reflect"
	"testing"
)

type judgeSquareSumTestPair struct {
	input int
	out   bool
}

var judgeSquareSumTestCases = []judgeSquareSumTestPair{
	{5, true},
	{3, false},
}

func TestEvalJudgeSquareSum(t *testing.T) {
	for _, pair := range judgeSquareSumTestCases {
		var newOut = judgeSquareSum(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
