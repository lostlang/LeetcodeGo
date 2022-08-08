package leetcode

import (
	"reflect"
	"testing"
)

type findComplementTestPair struct {
	input int
	out   int
}

var findComplementTestCases = []findComplementTestPair{
	{5, 2},
	{1, 0},
	{2, 1},
	{4, 3},
	{6, 1},
}

func TestEvalFindComplement(t *testing.T) {
	for _, pair := range findComplementTestCases {
		var newOut = findComplement(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
