package leetcode

import (
	"reflect"
	"testing"
)

type mostFrequentEvenTestPair struct {
	input []int
	out   int
}

var mostFrequentEvenTestCases = []mostFrequentEvenTestPair{
	{[]int{0, 1, 2, 2, 4, 4, 1}, 2},
	{[]int{4, 4, 4, 9, 2, 4}, 4},
	{[]int{29, 47, 21, 41, 13, 37, 25, 7}, -1},
}

func TestEvalMostFrequentEven(t *testing.T) {
	for _, pair := range mostFrequentEvenTestCases {
		var newOut = mostFrequentEven(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
