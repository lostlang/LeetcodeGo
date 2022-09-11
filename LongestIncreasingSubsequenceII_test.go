package leetcode

import (
	"reflect"
	"testing"
)

type lengthOfLIS2TestPair struct {
	inputNums []int
	inputK    int
	out       int
}

var lengthOfLIS2TestCases = []lengthOfLIS2TestPair{
	{[]int{4, 2, 1, 4, 3, 4, 5, 8, 15}, 3, 5},
	{[]int{7, 4, 5, 1, 8, 12, 4, 7}, 5, 4},
	{[]int{1, 5}, 1, 1},
	{[]int{1, 100, 500, 100000, 100000}, 100000, 4},
	{[]int{1, 19, 6, 2, 11, 13, 10}, 18, 4},
}

func TestEvalLengthOfLIS2(t *testing.T) {
	for _, pair := range lengthOfLIS2TestCases {
		var newOut = lengthOfLIS2(pair.inputNums, pair.inputK)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputNums, pair.inputK,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
