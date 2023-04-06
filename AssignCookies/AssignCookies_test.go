package assigncookies

import (
	"reflect"
	"testing"
)

type findContentChildrenTestPair struct {
	inputG []int
	inputS []int
	output int
}

var findContentChildrenTestCases = []findContentChildrenTestPair{
	{[]int{1, 2, 3}, []int{1, 1}, 1},
	{[]int{1, 2}, []int{1, 2, 3}, 2},
	{[]int{1, 2, 3}, []int{}, 0},
}

func TestEvalFindContentChildren(t *testing.T) {
	for _, pair := range findContentChildrenTestCases {
		newOutput := findContentChildren(pair.inputG, pair.inputS)
		if !reflect.DeepEqual(newOutput, pair.output) {
			t.Error(
				"For", pair.inputG, pair.inputS,
				"expected", pair.output,
				"got", newOutput,
			)
		}
	}
}
