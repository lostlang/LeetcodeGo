package leetcode

import (
	"reflect"
	"testing"
)

type missingNumberTestPair struct {
	input []int
	out   int
}

var missingNumberTestCases = []missingNumberTestPair{
	{[]int{3, 0, 1}, 2},
	{[]int{9, 6, 4, 2, 3, 5, 7, 0, 1}, 8},
	{[]int{0, 1}, 2},
}

func TestEvalMissingNumber(t *testing.T) {
	for _, pair := range missingNumberTestCases {
		var newOut = missingNumber(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
