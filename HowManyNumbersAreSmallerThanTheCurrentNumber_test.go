package leetcode

import (
	"reflect"
	"testing"
)

type smallerNumbersThanCurrentTestPair struct {
	input []int
	out   []int
}

var smallerNumbersThanCurrentTestCases = []smallerNumbersThanCurrentTestPair{
	{[]int{8, 1, 2, 2, 3}, []int{4, 0, 1, 1, 3}},
	{[]int{6, 5, 4, 8}, []int{2, 1, 0, 3}},
	{[]int{7, 7, 7, 7}, []int{0, 0, 0, 0}},
}

func TestEvalSmallerNumbersThanCurrent(t *testing.T) {
	for _, pair := range smallerNumbersThanCurrentTestCases {
		var newOut = smallerNumbersThanCurrent(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
