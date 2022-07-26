package leetcode

import (
	"reflect"
	"testing"
)

type nextGreaterElementTestPair struct {
	inputTarget []int
	inputArray  []int
	out         []int
}

var nextGreaterElementTestCases = []nextGreaterElementTestPair{
	{[]int{4, 1, 2}, []int{1, 3, 4, 2}, []int{-1, 3, -1}},
	{[]int{2, 4}, []int{1, 2, 3, 4}, []int{3, -1}},
}

func TestEvalNextGreaterElement(t *testing.T) {
	for _, pair := range nextGreaterElementTestCases {
		var newOut = nextGreaterElement(pair.inputTarget, pair.inputArray)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputTarget,
				"and", pair.inputArray,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
