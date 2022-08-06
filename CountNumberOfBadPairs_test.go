package leetcode

import (
	"reflect"
	"testing"
)

type countBadPairsTestPair struct {
	input []int
	out   int64
}

var countBadPairsTestCases = []countBadPairsTestPair{
	// {[]int{4, 1, 3, 3}, 5},
	// {[]int{1, 2, 3, 4, 5}, 0},
	// {[]int{43, 69, 66, 40, 33}, 10},
	// {[]int{28, 22, 74, 57, 47, 76, 2, 93, 6}, 36},
	// {[]int{1, 2, 2}, 2},
}

func TestEvalCountBadPairs(t *testing.T) {
	for _, pair := range countBadPairsTestCases {
		var newOut = countBadPairs(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
