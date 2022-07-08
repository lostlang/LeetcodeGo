package leetcode

import (
	"reflect"
	"testing"
)

type sumTestPair struct {
	inputA int
	inputB int
	out    int
}

var sumTestCases = []sumTestPair{
	{12, 5, 17},
	{-10, 4, -6},
}

func TestEvalsum(t *testing.T) {
	for _, pair := range sumTestCases {
		var newOut = sum(pair.inputA, pair.inputB)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputA,
				"and", pair.inputB,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
