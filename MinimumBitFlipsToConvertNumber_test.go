package leetcode

import (
	"reflect"
	"testing"
)

type minBitFlipsTestPair struct {
	inputStart int
	inputGoal  int
	out        int
}

var minBitFlipsTestCases = []minBitFlipsTestPair{
	{10, 7, 3},
	{3, 4, 3},
}

func TestEvalMinBitFlips(t *testing.T) {
	for _, pair := range minBitFlipsTestCases {
		newOut := minBitFlips(pair.inputStart, pair.inputGoal)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputStart, pair.inputGoal,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
