package leetcode

import (
	"reflect"
	"testing"
)

type minGroupsTestPair struct {
	input [][]int
	out   int
}

var minGroupsTestCases = []minGroupsTestPair{
	{[][]int{[]int{5, 10}, []int{6, 8}, []int{1, 5}, []int{2, 3}, []int{1, 10}}, 3},
	{[][]int{[]int{1, 3}, []int{5, 6}, []int{8, 10}, []int{11, 13}}, 1},
	{[][]int{[]int{1, 5}, []int{5, 8}}, 2},
}

func TestEvalMinGroups(t *testing.T) {
	for _, pair := range minGroupsTestCases {
		var newOut = minGroups(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
