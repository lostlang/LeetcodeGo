package leetcode

import (
	"reflect"
	"testing"
)

type spiralOrderTestPair struct {
	input [][]int
	out   []int
}

var spiralOrderTestCases = []spiralOrderTestPair{
	{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
	{[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}, []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},
}

func TestEvalSpiralOrder(t *testing.T) {
	for _, pair := range spiralOrderTestCases {
		var newOut = spiralOrder(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
