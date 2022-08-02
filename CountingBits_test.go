package leetcode

import (
	"reflect"
	"testing"
)

type countBitsTestPair struct {
	input int
	out   []int
}

var countBitsTestCases = []countBitsTestPair{
	{2, []int{0, 1, 1}},
	{5, []int{0, 1, 1, 2, 1, 2}},
}

func TestEvalCountBits(t *testing.T) {
	for _, pair := range countBitsTestCases {
		var newOut = countBits(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
