package leetcode

import (
	"reflect"
	"testing"
)

type countNegativesTestPair struct {
	input [][]int
	out   int
}

var countNegativesTestCases = []countNegativesTestPair{
	{[][]int{[]int{4, 3, 2, -1}, []int{3, 2, 1, -1}, []int{1, 1, -1, -2}, []int{-1, -1, -2, -3}}, 8},
	{[][]int{[]int{3, 2}, []int{1, 0}}, 0},
}

func TestEvalCountNegatives(t *testing.T) {
	for _, pair := range countNegativesTestCases {
		var newOut = countNegatives(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
