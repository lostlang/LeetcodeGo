package leetcode

import (
	"reflect"
	"testing"
)

type robTestPair struct {
	input []int
	out   int
}

var robTestCases = []robTestPair{
	// {[]int{1, 2, 3, 1}, 4},
	// {[]int{2, 7, 9, 3, 1}, 12},
	// {[]int{2, 1, 1, 2}, 4},
}

func TestEvalRob(t *testing.T) {
	for _, pair := range robTestCases {
		var newOut = rob(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
