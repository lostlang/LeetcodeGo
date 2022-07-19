package leetcode

import (
	"reflect"
	"testing"
)

type getRowTestPair struct {
	input int
	out   []int
}

var getRowTestCases = []getRowTestPair{
	{3, []int{1, 3, 3, 1}},
	{0, []int{1}},
	{1, []int{1, 1}},
}

func TestEvalGetRow(t *testing.T) {
	for _, pair := range getRowTestCases {
		var newOut = getRow(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
