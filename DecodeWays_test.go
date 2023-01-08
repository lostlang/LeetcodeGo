package leetcode

import (
	"reflect"
	"testing"
)

type numDecodingsTestPair struct {
	input string
	out   int
}

var numDecodingsTestCases = []numDecodingsTestPair{
	{"12", 2},
	{"226", 3},
	{"06", 0},
	{"10", 1},
	{"2101", 1},
	{"1234", 3},
	{"230", 0},
}

func TestEvalNumDecodings(t *testing.T) {
	for _, pair := range numDecodingsTestCases {
		var newOut = numDecodings(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
