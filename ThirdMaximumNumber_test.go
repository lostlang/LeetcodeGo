package leetcode

import (
	"reflect"
	"testing"
)

type thirdMaxTestPair struct {
	input []int
	out   int
}

var thirdMaxTestCases = []thirdMaxTestPair{
	{[]int{1, 2}, 2},
	{[]int{2, 2, 3, 1}, 1},
	{[]int{3, 2, 1}, 1},
	{[]int{2, 2, 2, 1}, 2},
	{[]int{1, 2, 2, 5, 3, 5}, 2},
}

func TestEvalThirdMax(t *testing.T) {
	for _, pair := range thirdMaxTestCases {
		var newOut = thirdMax(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
