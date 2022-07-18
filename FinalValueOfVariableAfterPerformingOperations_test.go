package leetcode

import (
	"reflect"
	"testing"
)

type finalValueAfterOperationsTestPair struct {
	input []string
	out   int
}

var finalValueAfterOperationsTestCases = []finalValueAfterOperationsTestPair{
	{[]string{"--X", "X++", "X++"}, 1},
	{[]string{"++X", "++X", "X++"}, 3},
	{[]string{"X++", "++X", "--X", "X--"}, 0},
}

func TestEvalFinalValueAfterOperations(t *testing.T) {
	for _, pair := range finalValueAfterOperationsTestCases {
		var newOut = finalValueAfterOperations(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
