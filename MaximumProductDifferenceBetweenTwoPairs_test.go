package leetcode

import (
	"reflect"
	"testing"
)

type maxProductDifferenceTestPair struct {
	input []int
	out   int
}

var maxProductDifferenceTestCases = []maxProductDifferenceTestPair{
	{[]int{5, 6, 2, 7, 4}, 34},
	{[]int{4, 2, 5, 9, 7, 4, 8}, 64},
}

func TestEvalMaxProductDifference(t *testing.T) {
	for _, pair := range maxProductDifferenceTestCases {
		var newOut = maxProductDifference(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
