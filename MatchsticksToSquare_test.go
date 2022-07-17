package leetcode

import (
	"reflect"
	"testing"
)

type makesquareTestPair struct {
	input []int
	out   bool
}

var makesquareTestCases = []makesquareTestPair{
	{[]int{1, 1, 2, 2, 2}, true},
	{[]int{3, 3, 3, 3, 4}, false},
	{[]int{3, 9, 2, 2, 2, 9, 10, 8, 3, 9, 10, 10, 1, 9, 9}, false},
	{[]int{5, 5, 5, 5, 4, 4, 4, 4, 3, 3, 3, 3}, true},
}

func TestEvalMakesquare(t *testing.T) {
	for _, pair := range makesquareTestCases {
		var newOut = makesquare(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
