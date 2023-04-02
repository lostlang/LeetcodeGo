package leetcode

import (
	"reflect"
	"testing"
)

type differenceOfSumTestPair struct {
	input []int
	out   int
}

var differenceOfSumTestCases = []differenceOfSumTestPair{
	{[]int{1, 2, 3, 4}, 0},
	{[]int{1, 15, 6, 3}, 9},
}

func TestEvalDifferenceOfSum(t *testing.T) {
	for _, pair := range differenceOfSumTestCases {
		newOut := differenceOfSum(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
