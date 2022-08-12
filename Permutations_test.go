package leetcode

import (
	"reflect"
	"testing"
)

type permuteTestPair struct {
	input []int
	out   [][]int
}

var permuteTestCases = []permuteTestPair{
	{[]int{1, 2, 3}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
}

func TestEvalPermute(t *testing.T) {
	for _, pair := range permuteTestCases {
		var newOut = permute(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
