package leetcode

import (
	"reflect"
	"testing"
)

type numIdenticalPairsTestPair struct {
	input []int
	out   int
}

var numIdenticalPairsTestCases = []numIdenticalPairsTestPair{
	{[]int{1, 2, 3, 1, 1, 3}, 4},
	{[]int{1, 1, 1, 1}, 6},
	{[]int{1, 2, 3}, 0},
}

func TestEvalNumIdenticalPairs(t *testing.T) {
	for _, pair := range numIdenticalPairsTestCases {
		var newOut = numIdenticalPairs(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
