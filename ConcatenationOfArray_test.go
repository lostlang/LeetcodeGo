package leetcode

import (
	"reflect"
	"testing"
)

type getConcatenationTestPair struct {
	input []int
	out   []int
}

var getConcatenationTestCases = []getConcatenationTestPair{
	{[]int{1, 2, 1}, []int{1, 2, 1, 1, 2, 1}},
	{[]int{1, 3, 2, 1}, []int{1, 3, 2, 1, 1, 3, 2, 1}},
}

func TestEvalGetConcatenation(t *testing.T) {
	for _, pair := range getConcatenationTestCases {
		var newOut = getConcatenation(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
