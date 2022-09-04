package leetcode

import (
	"reflect"
	"testing"
)

type numberOfWaysTestPair struct {
	inputStartPos int
	inputEndPos   int
	inputK        int
	out           int
}

var numberOfWaysTestCases = []numberOfWaysTestPair{
	{1, 2, 3, 3},
	{2, 5, 10, 0},
}

func TestEvalNumberOfWays(t *testing.T) {
	for _, pair := range numberOfWaysTestCases {
		var newOut = numberOfWays(pair.inputStartPos, pair.inputEndPos, pair.inputK)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputStartPos, pair.inputEndPos, pair.inputK,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
