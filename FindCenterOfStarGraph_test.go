package leetcode

import (
	"reflect"
	"testing"
)

type findCenterTestPair struct {
	input [][]int
	out   int
}

var findCenterTestCases = []findCenterTestPair{
	{[][]int{{1, 2}, {2, 3}, {4, 2}}, 2},
	{[][]int{{1, 2}, {5, 1}, {1, 3}, {1, 4}}, 1},
}

func TestEvalFindCenter(t *testing.T) {
	for _, pair := range findCenterTestCases {
		newOut := findCenter(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
