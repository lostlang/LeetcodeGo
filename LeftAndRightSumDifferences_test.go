package leetcode

import (
	"reflect"
	"testing"
)

type leftRigthDifferenceTestPair struct {
	input []int
	out   []int
}

var leftRigthDifferenceTestCases = []leftRigthDifferenceTestPair{
	{[]int{10, 4, 8, 3}, []int{15, 1, 11, 22}},
	{[]int{1}, []int{0}},
}

func TestEvalLeftRigthDifference(t *testing.T) {
	for _, pair := range leftRigthDifferenceTestCases {
		newOut := leftRigthDifference(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
