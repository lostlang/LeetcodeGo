package leetcode

import (
	"reflect"
	"testing"
)

type backspaceCompareTestPair struct {
	inputS string
	inputT string
	out    bool
}

var backspaceCompareTestCases = []backspaceCompareTestPair{
	{"ab#c", "ad#c", true},
	{"ab##", "c#d#", true},
	{"a#c", "b", false},
	{"a##c", "#a#c", true},
}

func TestEvalBackspaceCompare(t *testing.T) {
	for _, pair := range backspaceCompareTestCases {
		var newOut = backspaceCompare(pair.inputS, pair.inputT)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputS,
				"and", pair.inputT,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
