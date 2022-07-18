package leetcode

import (
	"reflect"
	"testing"
)

type shuffleTestPair struct {
	inputArray []int
	inputN     int
	out        []int
}

var shuffleTestCases = []shuffleTestPair{
	{[]int{2, 5, 1, 3, 4, 7}, 3, []int{2, 3, 5, 4, 1, 7}},
	{[]int{1, 2, 3, 4, 4, 3, 2, 1}, 4, []int{1, 4, 2, 3, 3, 2, 4, 1}},
	{[]int{1, 1, 2, 2}, 2, []int{1, 2, 1, 2}},
}

func TestEvalShuffle(t *testing.T) {
	for _, pair := range shuffleTestCases {
		var newOut = shuffle(pair.inputArray, pair.inputN)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputArray,
				"and", pair.inputN,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
