package AddStrings

import (
	"reflect"
	"testing"
)

type addStringsTestPair struct {
	inputS1 string
	inputS2 string
	output  string
}

var addStringsTestCases = []addStringsTestPair{
	{"11", "123", "134"},
	{"0", "0", "0"},
	{"456", "77", "533"},
}

func TestEvalAddStrings(t *testing.T) {
	for _, pair := range addStringsTestCases {
		newOutput := addStrings(pair.inputS1, pair.inputS2)
		if !reflect.DeepEqual(newOutput, pair.output) {
			t.Error(
				"For", pair.inputS1, pair.inputS2,
				"expected", pair.output,
				"got", newOutput,
			)
		}
	}
}
