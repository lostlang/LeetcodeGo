package leetcode

import (
	"reflect"
	"testing"
)

type largestPerimeterTestPair struct {
	input []int
	out   int
}

var largestPerimeterTestCases = []largestPerimeterTestPair{
	{[]int{2, 1, 2}, 5},
	{[]int{1, 1, 2}, 0},
}

func TestEvalLargestPerimeter(t *testing.T) {
	for _, pair := range largestPerimeterTestCases {
		var newOut = largestPerimeter(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
