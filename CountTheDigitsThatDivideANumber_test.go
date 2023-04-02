package leetcode

import (
	"reflect"
	"testing"
)

type countDigitsTestPair struct {
	input int
	out   int
}

var countDigitsTestCases = []countDigitsTestPair{
	{7, 1},
	{121, 2},
	{1248, 4},
}

func TestEvalCountDigits(t *testing.T) {
	for _, pair := range countDigitsTestCases {
		newOut := countDigits(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
