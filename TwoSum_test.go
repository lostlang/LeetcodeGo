package leetcode

import (
	"reflect"
	"testing"
)

type twoSumTestPair struct {
	input_n []int
	input_t int
	out     []int
}

var twoSumTestCases = []twoSumTestPair{
	{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	{[]int{3, 2, 4}, 6, []int{1, 2}},
	{[]int{3, 3}, 6, []int{0, 1}},
}

func TestEvalTwoSum(t *testing.T) {
	for _, pair := range twoSumTestCases {
		var newOut = twoSum(pair.input_n, pair.input_t)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input_n,
				"and", pair.input_t,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
