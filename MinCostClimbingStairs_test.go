package leetcode

import (
	"reflect"
	"testing"
)

type minCostClimbingStairsTestPair struct {
	input []int
	out   int
}

var minCostClimbingStairsTestCases = []minCostClimbingStairsTestPair{
	{[]int{10, 15, 20}, 15},
	{[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, 6},
}

func TestEvalMinCostClimbingStairs(t *testing.T) {
	for _, pair := range minCostClimbingStairsTestCases {
		var newOut = minCostClimbingStairs(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
