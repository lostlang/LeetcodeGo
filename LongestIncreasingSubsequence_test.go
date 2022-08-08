package leetcode

import (
	"reflect"
	"testing"
)

type lengthOfLISTestPair struct {
	input []int
	out   int
}

var lengthOfLISTestCases = []lengthOfLISTestPair{
	{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
	{[]int{7, 7, 7, 7}, 1},
}

func TestEvalLengthOfLIS(t *testing.T) {
	for _, pair := range lengthOfLISTestCases {
		var newOut = lengthOfLIS(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
