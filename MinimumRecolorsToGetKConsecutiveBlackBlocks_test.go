package leetcode

import (
	"reflect"
	"testing"
)

type minimumRecolorsTestPair struct {
	inputBlock string
	inputK     int
	out        int
}

var minimumRecolorsTestCases = []minimumRecolorsTestPair{
	// {"WBBWWBBWBW", 7, 3},
	// {"WBWBBBW", 2, 0},
	// {"BWWWBB", 6, 3},
	// {"BBBBBWWBBWBWBWWWBWBWBBBBWBBBBWBWBWBWBWWBWWBWBWWWWBBWWWWBWWWWBWBBWBBWBBWWW", 29, 10},
}

func TestEvalMinimumRecolors(t *testing.T) {
	for _, pair := range minimumRecolorsTestCases {
		var newOut = minimumRecolors(pair.inputBlock, pair.inputK)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputBlock,
				"and", pair.inputK,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
