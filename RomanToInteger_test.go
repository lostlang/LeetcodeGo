package leetcode

import (
	"reflect"
	"testing"
)

type romanToIntTestPair struct {
	input string
	out   int
}

var romanToIntTestCases = []romanToIntTestPair{
	{"III", 3},
	{"LVIII", 58},
	{"MCMXCIV", 1994},
}

func TestEvalRomanToInt(t *testing.T) {
	for _, pair := range romanToIntTestCases {
		var newOut = romanToInt(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
