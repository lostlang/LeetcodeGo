package leetcode

import (
	"reflect"
	"testing"
)

type lastStoneWeightTestPair struct {
	input []int
	out   int
}

var lastStoneWeightTestCases = []lastStoneWeightTestPair{
	{[]int{2, 7, 4, 1, 8, 1}, 1},
	{[]int{1}, 1},
}

func TestEvalLastStoneWeight(t *testing.T) {
	for _, pair := range lastStoneWeightTestCases {
		var newOut = lastStoneWeight(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
