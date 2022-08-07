package leetcode

import (
	"reflect"
	"testing"
)

type squareIsWhiteTestPair struct {
	input string
	out   bool
}

var squareIsWhiteTestCases = []squareIsWhiteTestPair{
	{"a1", false},
	{"h3", true},
	{"c7", false},
	{"c2", true},
}

func TestEvalSquareIsWhite(t *testing.T) {
	for _, pair := range squareIsWhiteTestCases {
		var newOut = squareIsWhite(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
