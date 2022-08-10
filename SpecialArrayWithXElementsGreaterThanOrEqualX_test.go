package leetcode

import (
	"reflect"
	"testing"
)

type specialArrayTestPair struct {
	input []int
	out   int
}

var specialArrayTestCases = []specialArrayTestPair{
	{[]int{3, 5}, 2},
	{[]int{0, 0}, -1},
	{[]int{0, 4, 3, 0, 4}, 3},
	{[]int{3, 6, 7, 7, 0}, -1},
	{[]int{100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100}, 100},
}

func TestEvalSpecialArray(t *testing.T) {
	for _, pair := range specialArrayTestCases {
		var newOut = specialArray(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
