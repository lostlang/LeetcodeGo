package leetcode

import (
	"reflect"
	"testing"
)

type findMiddleIndexTestPair struct {
	input []int
	out   int
}

var findMiddleIndexTestCases = []findMiddleIndexTestPair{
	{[]int{2, 3, -1, 8, 4}, 3},
	{[]int{1, -1, 4}, 2},
	{[]int{2, 5}, -1},
}

func TestEvalFindMiddleIndex(t *testing.T) {
	for _, pair := range findMiddleIndexTestCases {
		var newOut = findMiddleIndex(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
