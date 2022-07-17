package leetcode

import (
	"reflect"
	"testing"
)

type climbStairsTestPair struct {
	input int
	out   int
}

var climbStairsTestCases = []climbStairsTestPair{
	{2, 2},
	{3, 3},
}

func TestEvalClimbStairs(t *testing.T) {
	for _, pair := range climbStairsTestCases {
		var newOut = climbStairs(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
