package leetcode

import (
	"reflect"
	"testing"
)

type findKthPositiveTestPair struct {
	inputArr []int
	inputK   int
	out      int
}

var findKthPositiveTestCases = []findKthPositiveTestPair{
	{[]int{2, 3, 4, 7, 11}, 5, 9},
	{[]int{1, 2, 3, 4}, 2, 6},
	{[]int{19}, 2, 2},
}

func TestEvalFindKthPositive(t *testing.T) {
	for _, pair := range findKthPositiveTestCases {
		var newOut = findKthPositive(pair.inputArr, pair.inputK)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputArr,
				"and", pair.inputK,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
