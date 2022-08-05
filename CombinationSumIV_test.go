package leetcode

import (
	"reflect"
	"testing"
)

type combinationSum4TestPair struct {
	inputArr    []int
	inputTarget int
	out         int
}

var combinationSum4TestCases = []combinationSum4TestPair{
	{[]int{1, 2, 3}, 4, 7},
	{[]int{9}, 3, 0},
}

func TestEvalCombinationSum4(t *testing.T) {
	for _, pair := range combinationSum4TestCases {
		var newOut = combinationSum4(pair.inputArr, pair.inputTarget)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputArr,
				"and", pair.inputTarget,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
