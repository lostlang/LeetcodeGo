package leetcode

import (
	"reflect"
	"testing"
)

type checkStraightLineTestPair struct {
	input [][]int
	out   bool
}

var checkStraightLineTestCases = []checkStraightLineTestPair{
	{[][]int{[]int{1, 2}, []int{2, 3}, []int{3, 4}, []int{4, 5}, []int{5, 6}, []int{6, 7}}, true},
	{[][]int{[]int{1, 1}, []int{2, 2}, []int{3, 4}, []int{4, 5}, []int{5, 6}, []int{7, 7}}, false},
}

func TestEvalCheckStraightLine(t *testing.T) {
	for _, pair := range checkStraightLineTestCases {
		var newOut = checkStraightLine(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
