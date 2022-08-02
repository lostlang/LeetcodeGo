package leetcode

import (
	"reflect"
	"testing"
)

type plusOneTestPair struct {
	input []int
	out   []int
}

var plusOneTestCases = []plusOneTestPair{
	{[]int{1, 2, 3}, []int{1, 2, 4}},
	{[]int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
	{[]int{9}, []int{1, 0}},
}

func TestEvalPlusOne(t *testing.T) {
	for _, pair := range plusOneTestCases {
		var newOut = plusOne(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
