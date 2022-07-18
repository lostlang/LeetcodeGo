package leetcode

import (
	"reflect"
	"testing"
)

type maximumWealthTestPair struct {
	input [][]int
	out   int
}

var maximumWealthTestCases = []maximumWealthTestPair{
	{[][]int{[]int{1, 2, 3}, []int{1, 2, 3}}, 6},
	{[][]int{[]int{1, 5}, []int{7, 3}, []int{3, 5}}, 10},
	{[][]int{[]int{2, 8, 7}, []int{7, 1, 3}, []int{1, 9, 5}}, 17},
}

func TestEvalMaximumWealth(t *testing.T) {
	for _, pair := range maximumWealthTestCases {
		var newOut = maximumWealth(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
