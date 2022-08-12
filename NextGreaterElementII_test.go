package leetcode

import (
	"reflect"
	"testing"
)

type nextGreaterElementsTestPair struct {
	input []int
	out   []int
}

var nextGreaterElementsTestCases = []nextGreaterElementsTestPair{
	{[]int{1, 2, 1}, []int{2, -1, 2}},
	{[]int{1, 2, 3, 4, 5}, []int{2, 3, 4, 5, -1}},
}

func TestEvalNextGreaterElements(t *testing.T) {
	for _, pair := range nextGreaterElementsTestCases {
		var newOut = nextGreaterElements(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
