package leetcode

import (
	"reflect"
	"testing"
)

type findContentChildrenTestPair struct {
	inputG []int
	inputS []int
	out    int
}

var findContentChildrenTestCases = []findContentChildrenTestPair{
	{[]int{1, 2, 3}, []int{1, 1}, 1},
	{[]int{1, 2}, []int{1, 2, 3}, 2},
	{[]int{1, 2, 3}, []int{}, 0},
}

func TestEvalFindContentChildren(t *testing.T) {
	for _, pair := range findContentChildrenTestCases {
		var newOut = findContentChildren(pair.inputG, pair.inputS)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputG,
				"and", pair.inputS,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
