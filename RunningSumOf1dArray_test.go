package leetcode

import (
	"reflect"
	"testing"
)

type runningSumTestPair struct {
	input []int
	out   []int
}

var runningSumTestCases = []runningSumTestPair{
	{[]int{1, 2, 3, 4}, []int{1, 3, 6, 10}},
	{[]int{1, 1, 1, 1, 1}, []int{1, 2, 3, 4, 5}},
	{[]int{3, 1, 2, 10, 1}, []int{3, 4, 6, 16, 17}},
}

func TestEvalRunningSum(t *testing.T) {
	for _, pair := range runningSumTestCases {
		var newOut = runningSum(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
