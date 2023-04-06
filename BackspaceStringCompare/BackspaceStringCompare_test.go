package backspacestringcompare

import (
	"reflect"
	"testing"
)

type backspaceCompareTestPair struct {
	inputS string
	inputT string
	output bool
}

var backspaceCompareTestCases = []backspaceCompareTestPair{
	{"ab#c", "ad#c", true},
	{"ab##", "c#d#", true},
	{"a#c", "b", false},
	{"a##c", "#a#c", true},
}

func TestEvalBackspaceCompare(t *testing.T) {
	for _, pair := range backspaceCompareTestCases {
		newOutput := backspaceCompare(pair.inputS, pair.inputT)
		if !reflect.DeepEqual(newOutput, pair.output) {
			t.Error(
				"For", pair.inputS, pair.inputT,
				"expected", pair.output,
				"got", newOutput,
			)
		}
	}
}
