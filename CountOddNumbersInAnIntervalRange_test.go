package leetcode

import (
	"reflect"
	"testing"
)

type countOddsTestPair struct {
	inputStart int
	inputEnd   int
	out        int
}

var countOddsTestCases = []countOddsTestPair{
	{3, 7, 3},
	{8, 10, 1},
}

func TestEvalCountOdds(t *testing.T) {
	for _, pair := range countOddsTestCases {
		var newOut = countOdds(pair.inputStart, pair.inputEnd)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputStart,
				"and", pair.inputEnd,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
