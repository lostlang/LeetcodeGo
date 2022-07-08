package leetcode

import (
	"reflect"
	"testing"
)

type moveZeroesTestPair struct {
	input []int
	out   []int
}

var moveZeroesTestCases = []moveZeroesTestPair{
	{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
	{[]int{0}, []int{0}},
}

func TestEvalMoveZeroes(t *testing.T) {
	for _, pair := range moveZeroesTestCases {
		moveZeroes(pair.input)
		if !reflect.DeepEqual(pair.input, pair.out) {
			t.Error(
				"Got", pair.input,
				"expected", pair.out)
		}
	}
}
