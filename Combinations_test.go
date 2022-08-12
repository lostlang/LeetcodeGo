package leetcode

import (
	"reflect"
	"testing"
)

type combineTestPair struct {
	inputN int
	inputK int
	out    [][]int
}

var combineTestCases = []combineTestPair{
	{1, 1, [][]int{{1}}},
	{4, 2, [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}},
}

func TestEvalCombine(t *testing.T) {
	for _, pair := range combineTestCases {
		var newOut = combine(pair.inputN, pair.inputK)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputN,
				"and", pair.inputK,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
